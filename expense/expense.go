package expense

import (
	"database/sql"
	"net/http"
)

type Expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

type Logger struct {
	Handler http.Handler
}

type handler struct {
	DB *sql.DB
}

func NewExpenseHandler(db *sql.DB) *handler {
	return &handler{db}
}
