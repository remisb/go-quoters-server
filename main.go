package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/remisb/go-quoters-server/server"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// SrvConfig has server configuration settings.
type SrvConfig struct {
	Host            string
	Port            int
	Log             string
	DebugHost       string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func NewSrvConfig() SrvConfig {
	return SrvConfig{
		Host: "localhost",
		Port: 8080,
		Log: "./quoters.log",
	}
}

// Addr returns server address in the form of Host:Port localhost:8080.
func (sc SrvConfig) Addr() string {
	return sc.Host + ":" + strconv.Itoa(sc.Port)
}

// Config structure to store application configuration settings.
type Config struct {
	Server SrvConfig
	Db     DbConfig
	Args   Args
}

func main() {
	config := NewConfig()
	if err := initLogger(&config); err != nil {
		fmt.Printf("error on initLogger err: %s", err)
		os.Exit(1)
	}

	if err := startAPIServerAndWait(config); err != nil {
		Sugar.Errorf("error on starting api server, error :", err)
		os.Exit(1)
	}
}

func NewConfig() Config {
	return Config{
		Server: NewSrvConfig(),
		Db: DbConfig{
			Host: "localhost",
			Port: "3306",
			User: "root",
			Password: "",
			Name: "quoter",
		},
		Args: NewConfigArgs(os.Args),
	}
}

func startAPIServerAndWait(config Config) error {
	dbx, err := startDatabase(config.Db)
	if err != nil {
		return err
	}

	defer func() {
		Sugar.Infof("main : Database Stopping : %s", config.Db.Host)
		dbx.Close()
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	apiServer := startAPIServer(config, dbx, shutdown, serverErrors)
	return waitShutdown(config.Server, apiServer, serverErrors, shutdown)
}

func startAPIServer(cfg Config, dbx *sqlx.DB,
	shutdownChan chan os.Signal,
	serverErrors chan error) *http.Server {

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/info", server.InfoHandler)

	api := http.Server{
		Addr:         cfg.Server.Addr(),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start the service listening for requests.
	go func() {
		Sugar.Infof("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()
	return &api
}

func startDatabase(dbConf DbConfig) (*sqlx.DB, error) {
	Sugar.Infof("main : Started : Initializing database support")

	dbx, err := OpenDb(dbConf)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to db")
	}

	return dbx, nil
}

func waitShutdown(serverConf SrvConfig, apiServer *http.Server, serverErrors chan error, shutdown chan os.Signal) error {
	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		Sugar.Infof("main : %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), serverConf.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := apiServer.Shutdown(ctx)
		if err != nil {
			Sugar.Infof("main : Graceful shutdown did not complete in %v : %v", serverConf.ShutdownTimeout, err)
			err = apiServer.Close()
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}
	return nil
}