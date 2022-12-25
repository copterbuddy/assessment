package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) GetExpenseByIdHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	result := Expense{}
	row := h.DB.QueryRow("SELECT * FROM expenses WHERE id = $1", id)
	switch err := row.Scan(&result.ID, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags)); err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusBadRequest, Err{Message: "Not found youe expense"})
	}

	return c.JSON(http.StatusOK, result)
}
