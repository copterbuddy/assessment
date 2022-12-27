//go:build integration
// +build integration

package expense

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/copterbuddy/assessment/converter"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_it_Update_Success_Case(t *testing.T) {

	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgresql://root:root@db/go-example-db?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		h := NewExpenseHandler(db)

		e.PUT("/expenses/:id", h.UpdateExpenseHandler)
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
	testcase := "2"

	want := Expense{
		ID:     2,
		Title:  "apple smoothie",
		Amount: 89,
		Note:   "no discount",
		Tags:   []string{"beverage"},
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/expenses/"+testcase, serverPort), strings.NewReader(converter.ReqString(want)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
	client := http.Client{}

	//Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	var resStruct Expense
	json.Unmarshal(byteBody, &resStruct)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, 2, resStruct.ID)
		assert.Equal(t, want.Title, resStruct.Title)
		assert.Equal(t, want.Amount, resStruct.Amount)
		assert.Equal(t, want.Note, resStruct.Note)
		assert.Equal(t, want.Tags, resStruct.Tags)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
