package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) UpdateExpenseHandler(c echo.Context) error {

	id := c.Param("id")

	idDigit, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Info(err)
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	var e Expense
	err = c.Bind(&e)
	if err != nil {
		c.Logger().Info(err)
		return c.JSON(http.StatusBadRequest, Err{Message: "bad request"})
	}

	if e.Title == "" || e.Amount == 0 || e.Note == "" || e.Tags == nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "data incorrect"})
	}

	result, err := h.DB.Exec(`
	UPDATE expenses
	set title=$1,amount=$2,note=$3,tags=$4
	WHERE id=$5;
	`, e.Title, e.Amount, e.Note, pq.Array(e.Tags), id)
	if err != nil {
		c.Logger().Info(err.Error())
		return c.JSON(http.StatusInternalServerError, Err{Message: "internal server error please contact admin"})
	}
	rows, err := result.RowsAffected()
	if err != nil {
		c.Logger().Info(err.Error())
		return c.JSON(http.StatusInternalServerError, Err{Message: "internal server error please contact admin"})
	}
	if rows != 1 {
		c.Logger().Info("expected single row affected, got %d rows affected", rows)
		return c.JSON(http.StatusInternalServerError, Err{Message: "internal server error please contact admin"})
	}

	resp := Expense{
		ID:     idDigit,
		Title:  e.Title,
		Amount: e.Amount,
		Note:   e.Note,
		Tags:   e.Tags,
	}

	return c.JSON(http.StatusOK, resp)
}
