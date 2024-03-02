package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
)

func TestNewRegisterTodo(t *testing.T) {
	a := assert.New(t)
	title := "タイトル1"

	todo := entity.NewRegisterTodo(title)

	a.Equal(title, todo.Title)
}

func TestNewChangeTodo(t *testing.T) {
	a := assert.New(t)
	id := 1
	title := "タイトル1"

	todo := entity.NewChangeTodo(id, title)

	a.Equal(id, todo.ID)
	a.Equal(title, todo.Title)
}
