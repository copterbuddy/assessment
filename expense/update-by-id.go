package expense

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
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

	fmt.Println("title is", e.Title, "amount is", e.Amount)

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
		ID:     e.ID,
		Title:  e.Title,
		Amount: e.Amount,
		Note:   e.Note,
		Tags:   e.Tags,
	}

	return c.JSON(http.StatusOK, resp)
}
