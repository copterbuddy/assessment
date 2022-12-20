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
	e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)

	c, rec := request(http.MethodGet, "/", strings.NewReader(""), e)

	h := handler{}
	err := h.Greeting(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

func request(method, url string, body io.Reader, e *echo.Echo) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	return c, rec
}

func Test_Create_When_No_Request_Body(t *testing.T) {

	var err error
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	h := handler{nil}
	c := e.NewContext(req, rec)

	err = h.CreateExpenseHandler(c)
	expected := "{\"message\":\"data incurrect\"}"

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
