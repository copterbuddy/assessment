package expense

import (
	"net/http"
	"testing"

	"github.com/copterbuddy/assessment/converter"
	"github.com/copterbuddy/assessment/request"
	"github.com/stretchr/testify/assert"
)

func Test_Update_Success(t *testing.T) {
	//Arrange
	testcase := Expense{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	want := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 89,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	ctx, rec := request.Request(http.MethodPut, request.Uri("expenses"), converter.ReqString(testcase))
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	h := handler{nil}

	//Act
	err := h.UpdateExpenseHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	ResponseBody := Expense{}
	converter.ResStruct(rec, &ResponseBody)

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, ResponseBody)
	}
}

func Test_Update_No_ID(t *testing.T) {
	//Arrange

	ctx, rec := request.Request(http.MethodPut, request.Uri("expenses"), "")
	h := handler{nil}

	//Act
	err := h.UpdateExpenseHandler(ctx)
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

func Test_Update_No_Req_Body(t *testing.T) {

	testcase := []Expense{
		{Title: "", Amount: 79, Note: "night market promotion discount 10 bath", Tags: []string{"food", "beverage"}},
		{Title: "strawberry smoothie", Amount: 0, Note: "night market promotion discount 10 bath", Tags: []string{"food", "beverage"}},
		{Title: "strawberry smoothie", Amount: 79, Note: "", Tags: nil},
		{Title: "", Amount: 0, Note: "", Tags: nil},
	}

	for _, c := range testcase {
		t.Run("invalid parameter", func(t *testing.T) {
			//Arrange
			ctx, rec := request.Request(http.MethodPut, request.Uri("expenses"), converter.ReqString(c))
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			h := handler{nil}

			//Act
			err := h.UpdateExpenseHandler(ctx)
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
		})
	}

}

func Test_Update_Bind_Req_Body_Err(t *testing.T) {

	testcase := []string{
		"{\"title\":10,\"amount\":79,\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":\"79\",\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":\"79\",\"note\":3,\"tags\":Tags:[]string{\"food\",\"beverage\"}}",
		"{\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"nightmarketpromotiondiscount10bath\",\"tags\":10}",
	}
	for _, c := range testcase {
		t.Run("invalid parameter", func(t *testing.T) {
			//Arrange
			ctx, rec := request.Request(http.MethodPut, request.Uri("expenses"), c)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			h := handler{nil}

			//Act
			err := h.UpdateExpenseHandler(ctx)
			if err != nil {
				t.Errorf("Test failed: %v", err)
			}

			ResponseBody := Err{}
			converter.ResStruct(rec, &ResponseBody)

			//Assert
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.Equal(t, Err{Message: "bad request"}, ResponseBody)
			}
		})
	}

}
