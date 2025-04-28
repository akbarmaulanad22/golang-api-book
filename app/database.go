package app

import (
	"api_book/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_api_book")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	
	return db

}