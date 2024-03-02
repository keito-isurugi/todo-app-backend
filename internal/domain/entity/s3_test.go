package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
)

func TestNewPutObjectInput(t *testing.T) {
	a := assert.New(t)
	poi := entity.NewPutObjectInput("key", "contentType", []byte("fileContent"))
	a.Equal("key", poi.Key)
	a.Equal("contentType", poi.ContentType)
	a.Equal([]byte("fileContent"), poi.FileContent)
}
