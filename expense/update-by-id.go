package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) UpdateExpenseHandler(c echo.Context) error {

	return c.JSON(http.StatusCreated, "ok")
}
