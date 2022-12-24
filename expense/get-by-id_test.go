//go:build unit
// +build unit

package expense

import (
	"net/http"
	"testing"

	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_Route_Get_Expense_By_Id(t *testing.T) {
	c, rec := request.Request(http.MethodGet, request.Uri("expenses", "1"), "")

	h := handler{}
	err := h.GetExpenseByIdHandler(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
