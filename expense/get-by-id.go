package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetExpenseByIdHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incurrect"})
	}

	return c.JSON(http.StatusOK, "ok")
}
