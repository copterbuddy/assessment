package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Expense_Handler_With_DBB(t *testing.T) {
	db, _, _ := sqlmock.New()
	newHandler := NewExpenseHandler(db)
	assert.NotNil(t, newHandler)
}
