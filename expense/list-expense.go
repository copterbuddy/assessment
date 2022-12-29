package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) ListExpenseHandler(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM expenses")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Not found your expense"})
	}
	defer rows.Close()

	result := []Expense{}
	item := Expense{}

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Title, &item.Amount, &item.Note, pq.Array(&item.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{
				Message: "internal server error please contact admin",
			})
		}
		result = append(result, item)
	}

	return c.JSON(http.StatusOK, result)
}
