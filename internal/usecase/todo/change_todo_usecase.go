package usecase

import (
	"context"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
)

type ChangeTodoUsecase interface {
	Exec(ctx context.Context, dh *entity.Todo) error
}

type changeTodoUsecaseImpl struct {
	todoRepo domain.Todo
}

func NewChangeTodoUsecase(todoRepo domain.Todo) ChangeTodoUsecase {
	return &changeTodoUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (c *changeTodoUsecaseImpl) Exec(ctx context.Context, todo *entity.Todo) error {
	err := c.todoRepo.ChangeTodo(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}
