package expense

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	//Arrange
	c, res := request(http.MethodGet, uri(""), strings.NewReader(""))
	h := handler{}

	//Act
	err := h.Greeting(c)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Hello, World!", res.Body.String())
	}
}

func Test_Create_When_No_Request_Body(t *testing.T) {
	//Arrange
	c, res := request(http.MethodGet, uri("expenses"), strings.NewReader(""))
	expected := Err{
		Message: "data incurrect",
	}
	h := handler{nil}

	//Act
	err := h.CreateExpenseHandler(c)

	ResponseBody := Err{}
	if res.Body == nil {
		panic(err)
	}
	json.Unmarshal([]byte(res.Body.Bytes()), &ResponseBody)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, ResponseBody)
	}
}

func request(method, url string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	return c, rec
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}
