package expense

import (
	"net/http"
	"testing"

	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_Update_Route_Success(t *testing.T) {
	//Arrange

	ctx, res := request.Request(http.MethodPut, request.Uri("expenses"), "")
	h := handler{nil}

	//Act
	err := h.UpdateExpenseHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, res.Code)
	}
}
