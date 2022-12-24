//go:build integration
// +build integration

package expense

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const serverPort = 2565

func Test_it_Create_Success_Case(t *testing.T) {

	eh := echo.New()
	go func(e *echo.Echo) {
		h := NewExpenseHandler(nil)

		e.POST("/expenses", h.CreateExpenseHandler)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

	//Arrange
	// testcase := Expense{
	// 	ID:     0,
	// 	Title:  "strawberry smoothie",
	// 	Amount: 79,
	// 	Note:   "night market promotion discount 10 bath",
	// 	Tags:   []string{"food", "beverage"},
	// }
	// c, res := request.Request(http.MethodPost, request.Uri("expenses"), converter.ReqString(testcase))
	// h := handler{db}
	reqBody := `
	{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}
	`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	//Act
	// err := h.Greeting(c)
	// if err != nil {
	// 	t.Errorf("Test failed: %v", err)
	// }
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	//Assert
	// assert.Equal(t, http.StatusOK, res.Code)
	// assert.Equal(t, "Hello, World!", res.Body.String())
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		// assert.Equal(t, "Hello, World!", string(byteBody))
		log.Println(fmt.Sprintf(fmt.Sprintf("int result is : %v", string(byteBody))))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
