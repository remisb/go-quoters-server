package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Config struct is used to to store db connection settings.
type DbConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	Name             string `yaml:"name"`
	DisableTLS       bool   `yaml:"disableTls"`
	connectionString string
}

func InitDb(dbCfg *DbConfig) {
	dbCfg.connectionString = mysqlConnectString(*dbCfg)
}

func (db DbConfig) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", db.connectionString)
}

func mysqlConnectString(config DbConfig) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
}
