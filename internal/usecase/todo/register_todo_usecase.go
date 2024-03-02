package usecase

import (
	"context"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
)

type RegisterTodoUsecase interface {
	Exec(ctx context.Context, todo *entity.Todo) (int, error)
}

type registerTodoUsecaseImpl struct {
	todoRepo domain.Todo
}

func NewRegisterTodoUsecase(todoRepo domain.Todo) RegisterTodoUsecase {
	return &registerTodoUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (r *registerTodoUsecaseImpl) Exec(ctx context.Context, todo *entity.Todo) (int, error) {
	todoID, err := r.todoRepo.RegisterTodo(ctx, todo)
	if err != nil {
		return 0, err
	}
	return todoID, nil
}
