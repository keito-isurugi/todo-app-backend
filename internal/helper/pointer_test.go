package helper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/helper"
)

func TestBoolToPtr(t *testing.T) {
	b := true
	ptr := helper.BoolToPtr(b)
	assert.NotNil(t, ptr)
	assert.Equal(t, b, *ptr)
}

func TestStringToPtr(t *testing.T) {
	s := "hello"
	ptr := helper.StringToPtr(s)
	assert.NotNil(t, ptr)
	assert.Equal(t, s, *ptr)
}

func TestIntToPtr(t *testing.T) {
	i := 42
	ptr := helper.IntToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestInt8ToPtr(t *testing.T) {
	i := int8(42)
	ptr := helper.Int8ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestInt32ToPtr(t *testing.T) {
	i := int32(42)
	ptr := helper.Int32ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestInt64ToPtr(t *testing.T) {
	i := int64(42)
	ptr := helper.Int64ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestUintToPtr(t *testing.T) {
	i := uint(42)
	ptr := helper.UintToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestUint8ToPtr(t *testing.T) {
	i := uint8(42)
	ptr := helper.Uint8ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestUint32ToPtr(t *testing.T) {
	i := uint32(42)
	ptr := helper.Uint32ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestUint64ToPtr(t *testing.T) {
	i := uint64(42)
	ptr := helper.Uint64ToPtr(i)
	assert.NotNil(t, ptr)
	assert.Equal(t, i, *ptr)
}

func TestTimeToPtr(t *testing.T) {
	now := time.Now()
	ptr := helper.TimeToPtr(now)
	assert.NotNil(t, ptr)
	assert.Equal(t, now, *ptr)
}
