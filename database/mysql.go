package database

import (
	"cake-store-golang-restfull-api/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cake-store")
	helper.IfError(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
