package usecase

import (
	"context"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
)

type ListTodosUsecase interface {
	Exec(ctx context.Context) (entity.ListTodos, error)
}

type listTodosUsecaseImpl struct {
	todoRepo domain.Todo
}

func NewListTodosUsecase(todoRepo domain.Todo) ListTodosUsecase {
	return &listTodosUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (g *listTodosUsecaseImpl) Exec(ctx context.Context) (entity.ListTodos, error) {
	todos, err := g.todoRepo.ListTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
