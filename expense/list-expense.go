package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) ListExpenseHandler(c echo.Context) error {

	result := []Expense{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
	}

	return c.JSON(http.StatusOK, result)
}
