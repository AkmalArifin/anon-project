package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	tableName := os.Getenv("DB_TABLE")
	dataSources := username + ":" + password + "@tcp(localhost:3306)/" + tableName

	var err error
	DB, err = sql.Open("mysql", dataSources)

	if err != nil {
		panic("could not connect database")
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
