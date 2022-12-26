package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) UpdateExpenseHandler(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	var e Expense
	err := c.Bind(&e)
	if err != nil {
		c.Logger().Info(err)
		return c.JSON(http.StatusBadRequest, Err{Message: "bad request"})
	}

	if e.Title == "" || e.Amount == 0 || e.Note == "" || e.Tags == nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	result := Expense{
		ID:     1,
		Title:  e.Title,
		Amount: 89,
		Note:   e.Note,
		Tags:   e.Tags,
	}

	return c.JSON(http.StatusOK, result)
}
