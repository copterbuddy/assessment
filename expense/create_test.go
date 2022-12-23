package expense

import (
	"net/http"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	//Arrange
	c, res := request.Request(http.MethodGet, request.Uri(""), converter.ReqString(""))
	h := handler{}

	//Act
	err := h.Greeting(c)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	//Assert
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Hello, World!", res.Body.String())
}

func Test_Create_Success_Case(t *testing.T) {
	//Arrange
	testcase := Expense{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	want := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	ctx, res := request.Request(http.MethodPost, request.Uri("expenses"), converter.ReqString(testcase))
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("INSERT INTO expenses (.+) RETURNING id").
		WithArgs(testcase.Title, testcase.Amount, testcase.Note, `{"`+strings.Join(testcase.Tags, `","`)+`"}`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.CreateExpenseHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	ResponseBody := Expense{}
	converter.ResStruct(res, &ResponseBody)

	//Assert
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, want, ResponseBody)
}

func Test_Create_When_Query_Error(t *testing.T) {
	//Arrange
	testcase := Expense{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	ctx, res := request.Request(http.MethodPost, request.Uri("expenses"), converter.ReqString(testcase))
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("INSERT INTO expenses (.+) RETURNING id").
		WithArgs(testcase.Title, testcase.Amount, testcase.Note, testcase.Tags).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.CreateExpenseHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	ResponseBody := Err{}
	converter.ResStruct(res, &ResponseBody)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.NotEqual(t, "", ResponseBody)
}

func Test_Create_When_Request_Body_Bind_Error(t *testing.T) {
	testcase := []string{
		"{\"title\":10,\"amount\":79,\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":\"79\",\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":\"79\",\"note\":3,\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":10}",
	}

	for _, c := range testcase {
		t.Run("bind body error", func(t *testing.T) {
			//Arrange
			want := Err{
				Message: "bad request",
			}

			ctx, res := request.Request(http.MethodPost, request.Uri("expenses"), c)
			h := handler{nil}

			//Act
			err := h.CreateExpenseHandler(ctx)
			if err != nil {
				t.Errorf("Test failed: %v", err)
			}

			ResponseBody := Err{}
			converter.ResStruct(res, &ResponseBody)

			//Assert
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, want, ResponseBody)

		})
	}

}

func Test_Create_When_No_Request_Body(t *testing.T) {
	testcase := []Expense{
		{Title: "", Amount: 79, Note: "night market promotion discount 10 bath", Tags: []string{"food", "beverage"}},
		{Title: "strawberry smoothie", Amount: 0, Note: "night market promotion discount 10 bath", Tags: []string{"food", "beverage"}},
		{Title: "strawberry smoothie", Amount: 79, Note: "", Tags: nil},
		{Title: "", Amount: 0, Note: "", Tags: nil},
	}

	for _, c := range testcase {
		t.Run("invalid parameter", func(t *testing.T) {
			//Arrange
			want := Err{
				Message: "data incurrect",
			}

			ctx, res := request.Request(http.MethodPost, request.Uri("expenses"), converter.ReqString(c))
			h := handler{nil}

			//Act
			err := h.CreateExpenseHandler(ctx)
			if err != nil {
				t.Errorf("Test failed: %v", err)
			}

			ResponseBody := Err{}
			converter.ResStruct(res, &ResponseBody)

			//Assert
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, want, ResponseBody)

		})
	}

}
