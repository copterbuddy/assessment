package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) CreateExpenseHandler(c echo.Context) error {

	var e Expense
	err := c.Bind(&e)
	if err != nil {
		c.Logger().Info(err)
		return c.JSON(http.StatusBadRequest, Err{Message: "bad request"})
	}

	if e.Title == "" || e.Amount == 0 || e.Note == "" || e.Tags == nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incurrect"})
	}

	row := h.DB.QueryRow(`
	INSERT INTO expenses (title,amount,note,tags) 
	values ($1,$2,$3,$4) RETURNING id
	`,
		e.Title, e.Amount, e.Note, pq.Array(e.Tags))

	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, e)
}
