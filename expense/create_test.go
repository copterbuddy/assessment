package expense

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/copterbuddy/assessment/converter"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	//Arrange
	c, res := Request(http.MethodGet, Uri(""), converter.ReqString(""))
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

func Test_Create_When_No_Request_Body(t *testing.T) {
	//Arrange
	expense := Expense{
		ID:     0,
		Title:  "",
		Amount: 0,
		Note:   "",
		Tags:   nil,
	}
	c, res := Request(http.MethodGet, Uri("expenses"), converter.ReqString((expense)))
	expected := Err{
		Message: "data incurrect",
	}
	h := handler{nil}

	//Act
	err := h.CreateExpenseHandler(c)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	ResponseBody := Err{}
	converter.ResStruct(res, &ResponseBody)

	//Assert
	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.Equal(t, expected, ResponseBody)
}

func Request(method, url string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	return c, rec
}

func Uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}
