package models

import (
	"database/sql"
	"fmt"

	"github.com/Telinga-Digital/go-structure/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	driver = config.GetDBConfig().Connection
	user   = config.GetDBConfig().Username
	pass   = config.GetDBConfig().Password
	host   = config.GetDBConfig().Host
	port   = config.GetDBConfig().Port
	data   = config.GetDBConfig().Database
)

func MakeConnection() {
	switch driver {
	case "mysql":
		DB, _ = connectToMySQL()
	case "postgres":
		DB, _ = connectToPostgres()
	default:
		panic("Invalid database driver")
	}
}

func connectToMySQL() (*gorm.DB, error) {
	sqlDB, _ := sql.Open("mysql", fmt.Sprintf("%s%s@tcp(%s:%s)/%s", user, ":"+pass, host, port, data))

	return gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
}

func connectToPostgres() (*gorm.DB, error) {
	sqlDB, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, data))

	return gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
}
