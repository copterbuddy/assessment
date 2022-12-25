package expense

import (
	"net/http"
	"testing"

	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_Update_Route_Success(t *testing.T) {
	//Arrange

	c, res := request.Request(http.MethodPut, request.Uri("expenses"), "")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := handler{nil}

	//Act
	err := h.UpdateExpenseHandler(c)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, res.Code)
	}
}

func Test_Update_No_ID(t *testing.T) {
	//Arrange

	c, rec := request.Request(http.MethodPut, request.Uri("expenses"), "")
	h := handler{nil}

	//Act
	err := h.UpdateExpenseHandler(c)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	ResponseBody := Err{}
	converter.ResStruct(rec, &ResponseBody)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, Err{Message: "data incorrect"}, ResponseBody)
	}
}
