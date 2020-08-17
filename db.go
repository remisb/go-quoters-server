package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Config struct is used to to store db connection settings.
type DbConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	DisableTLS bool
}

func OpenDb(dbCfg DbConfig) (*sqlx.DB, error) {
	// Define SSL mode.
	//sslMode := "require"
	//if cfg.DisableTLS {
	//	sslMode = "disable"
	//}

	// Query parameters.
	//q := make(url.Values)
	//q.Set("sslmode", sslMode)
	//q.Set("timezone", "utc")

	// Construct url.
	//u := url.URL{
	//	Scheme:   "mysql",
	//	User:     url.UserPassword(dbCfg.User, dbCfg.Password),
	//	Host:     dbCfg.Host,
	//	Path:     dbCfg.Name,
	//	RawQuery: q.Encode(),
	//}

	//s := u.String()
	s := mysqlConnectString(dbCfg)
	//s :=  "root:@(localhost:3306)/quoter"
	return sqlx.Open("mysql", s)
}

func mysqlConnectString(config DbConfig) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
}
