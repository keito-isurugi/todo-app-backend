package helper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/helper"
)

func TestIsNilOrEmpty(t *testing.T) {
	a := assert.New(t)

	// Test nil
	var nilPtr *string
	a.True(helper.IsNilOrEmpty(nilPtr))
	a.True(helper.IsNilOrEmpty(nil))

	// Test empty string
	emptyStr := ""
	a.True(helper.IsNilOrEmpty(emptyStr))

	// Test non-empty string
	nonEmptyStr := "abc"
	a.False(helper.IsNilOrEmpty(nonEmptyStr))

	// Test zero time
	zeroTime := time.Time{}
	a.True(helper.IsNilOrEmpty(zeroTime))

	// Test non-zero time
	nonZeroTime := time.Now()
	a.False(helper.IsNilOrEmpty(nonZeroTime))

	// Test empty slice
	emptySlice := []int{}
	a.True(helper.IsNilOrEmpty(emptySlice))

	// Test non-empty slice
	nonEmptySlice := []int{1, 2, 3}
	a.False(helper.IsNilOrEmpty(nonEmptySlice))

	// Test boolean
	a.False(helper.IsNilOrEmpty(true))
	a.False(helper.IsNilOrEmpty(false))

	// Test integer
	a.True(helper.IsNilOrEmpty(0))
	a.False(helper.IsNilOrEmpty(1))

	// Test float
	a.False(helper.IsNilOrEmpty(0.0))
	a.False(helper.IsNilOrEmpty(1.1))
}
