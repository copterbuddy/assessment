package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
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

type Response struct {
	*http.Response
	err error
}

type handler struct {
	DB *sql.DB
}

func NewExpenseHandler(db *sql.DB) *handler {
	return &handler{db}
}

func (h *handler) Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
