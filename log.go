package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

// Logger has zap's Logger
var Logger *zap.Logger
// Sugar has zap's SugaredLogger
var Sugar *zap.SugaredLogger

func initLogger(config *Config) error {
	var logPath string
	var err error

	logCfg := zap.NewDevelopmentConfig()

	if !filepath.IsAbs(config.Server.Log) {
		logPath, err = filepath.Abs(config.Server.Log)
		if err != nil {
			return errors.Wrap(err, "unable to get absolute logfile path")
		}
	}

	fmt.Printf("Log file path: %s", logPath)

	logCfg.OutputPaths = append(logCfg.OutputPaths, logPath)
	Logger, err = logCfg.Build()

	if err != nil {
		err := fmt.Errorf("error in log init err: %v", err)
		fmt.Print(err)
	}
	Sugar = Logger.Sugar()
	logCfg.Level.SetLevel(zapcore.DebugLevel)

	Sugar.Infof("Logging is set to console and file: %s", logPath)
	Sugar.Sync()

	return nil
}
