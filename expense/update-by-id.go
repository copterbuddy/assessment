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

	return c.JSON(http.StatusCreated, "ok")
}
