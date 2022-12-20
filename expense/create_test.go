package expense

import (
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
	c, res := request(http.MethodGet, uri(""), strings.NewReader(""))

	h := handler{}
	err := h.Greeting(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Hello, World!", res.Body.String())
	}
}

func Test_Create_When_No_Request_Body(t *testing.T) {
	c, res := request(http.MethodGet, uri("expenses"), strings.NewReader(""))

	h := handler{nil}
	err := h.CreateExpenseHandler(c)
	expected := "{\"message\":\"data incurrect\"}"

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, strings.TrimSpace(res.Body.String()))
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
