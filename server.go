package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/copterbuddy/assessment/expense"
	"github.com/copterbuddy/assessment/request"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	SetupApi(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))

	go func() {
		if err := e.Start(":2565"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	GracefulShutdown(e)
}

func GracefulShutdown(e *echo.Echo) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func SetupApi(db *sql.DB, e *echo.Echo) {
	expenseHandler := expense.NewExpenseHandler(db)
	g := e.Group("/expenses")
	{
		g.Use(middleware.Logger())
		g.Use(middleware.Recover())
		g.Use(request.Auth)

		g.POST("", expenseHandler.CreateExpenseHandler)
		g.GET("/:id", expenseHandler.GetExpenseByIdHandler)
		g.PUT("/:id", expenseHandler.UpdateExpenseHandler)
		g.GET("", expenseHandler.ListExpenseHandler)
	}
}

func InitDB() (*sql.DB, error) {
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

	return db, nil
}
