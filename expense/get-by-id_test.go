//go:build unit
// +build unit

package expense

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Expense_By_Id(t *testing.T) {
	//Arrange
	want := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	c, rec := request.Request(http.MethodGet, request.Uri("expenses"), "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(want.ID, want.Title, want.Amount, want.Note, pq.Array(want.Tags))

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.GetExpenseByIdHandler(c)

	ResponseBody := Expense{}
	converter.ResStruct(rec, &ResponseBody)

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, ResponseBody)
	}
}

func Test_Get_Expense_By_Id_No_Param(t *testing.T) {
	//Arrange
	c, rec := request.Request(http.MethodGet, request.Uri("expenses"), "")

	//Act
	h := handler{}
	err := h.GetExpenseByIdHandler(c)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
