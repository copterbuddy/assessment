package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	InitDB()

	e := echo.New()

	e.Logger.Fatal(e.Start(":2565"))
}

var db *sql.DB

func InitDB() {
	url := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connection to database error ", url)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create database", err)
	}
}
