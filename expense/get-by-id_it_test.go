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

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_it_Get_Success_Case(t *testing.T) {

	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgresql://root:root@db/go-example-db?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		h := NewExpenseHandler(db)

		e.GET("/expenses/:id", h.GetExpenseByIdHandler)
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
	testcase := "1"

	want := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/"+testcase, serverPort), strings.NewReader(""))
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
	err = json.Unmarshal(byteBody, &resStruct)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEqual(t, 0, resStruct.ID)
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

func Test_it_Get_Success_Case_Not_Found(t *testing.T) {

	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgresql://root:root@db/go-example-db?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		h := NewExpenseHandler(db)

		e.GET("/expenses/:id", h.GetExpenseByIdHandler)
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
	testcase := "0"

	want := Err{
		Message: "Not found your expense",
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/"+testcase, serverPort), strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	//Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	var resStruct Err
	err = json.Unmarshal(byteBody, &resStruct)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NotEmpty(t, want, resStruct)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
