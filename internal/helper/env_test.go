package helper_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/helper"
)

func TestIsLocal(t *testing.T) {
	t.Run("ENV=local", func(t *testing.T) {
		a := assert.New(t)
		a.NoError(os.Setenv("ENV", "local"))
		isLocal := helper.IsLocal()
		a.True(isLocal)
	})

	t.Run("ENV=production", func(t *testing.T) {
		a := assert.New(t)
		a.NoError(os.Setenv("ENV", "production"))
		isLocal := helper.IsLocal()
		a.False(isLocal)
	})
}
