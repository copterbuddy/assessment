//go:build integration
// +build integration

package greeting

import (
	"net/http"
	"testing"

	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_it_GetGreeting(t *testing.T) {
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
