package expense

import (
	"net/http"
	"testing"

	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Expense_By_Id(t *testing.T) {
	//Arrange
	c, rec := request.Request(http.MethodGet, request.Uri("expenses"), "")
	h := handler{nil}

	//Act
	err := h.ListExpenseHandler(c)

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
