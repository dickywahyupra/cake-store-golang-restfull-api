package database

import (
	"cake-store-golang-restfull-api/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() *sql.DB {
	drive := helper.Env("DB_DRIVE")
	source := helper.Env("DB_USERNAME") + ":" + helper.Env("DB_PASSWORD") + "@tcp(" + helper.Env("DB_HOST") + ":" + helper.Env("DB_PORT") + ")/" + helper.Env("DB_NAME")

	db, err := sql.Open(drive, source)
	helper.IfError(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
