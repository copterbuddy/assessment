//go:build unit
// +build unit

package expense

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_List_Expense(t *testing.T) {
	//Arrange
	want := []Expense{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
	}

	c, rec := request.Request(http.MethodGet, request.Uri("expenses"), "")

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"})
	for _, item := range want {
		newsMockRows.AddRow(item.ID, item.Title, item.Amount, item.Note, pq.Array(item.Tags))
	}

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.ListExpenseHandler(c)
	ResponseBody := []Expense{}
	err = converter.ResStruct(rec, &ResponseBody)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, ResponseBody)
	}
}

func Test_List_Expense_Error(t *testing.T) {
	//Arrange
	want := Err{
		Message: "Not found your expense",
	}

	c, rec := request.Request(http.MethodGet, request.Uri("expenses"), "")

	newsMockRows := sql.ErrNoRows
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnError(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.ListExpenseHandler(c)
	ResponseBody := Err{}
	err = converter.ResStruct(rec, &ResponseBody)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, want, ResponseBody)
	}
}
