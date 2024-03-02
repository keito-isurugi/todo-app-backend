//go:generate mockgen -source=todo.go -destination=./mock/todo_mock.go
package domain

import (
	"context"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
)

type Todo interface {
	GetTodo(ctx context.Context, id int) (*entity.Todo, error)
	ListTodos(ctx context.Context) (entity.ListTodos, error)
	RegisterTodo(ctx context.Context, todo *entity.Todo) (int, error)
	ChangeTodoDoneFlag(ctx context.Context, todo *entity.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}
