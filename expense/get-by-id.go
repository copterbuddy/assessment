package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetExpenseByIdHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	result := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	return c.JSON(http.StatusOK, result)
}
