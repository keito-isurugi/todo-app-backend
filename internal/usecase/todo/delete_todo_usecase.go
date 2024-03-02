package usecase

import (
	"context"

	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
)

type DeleteTodoUsecase interface {
	Exec(ctx context.Context, id int) error
}

type deleteTodoUsecaseImpl struct {
	todoRepo domain.Todo
}

func NewDeleteTodoUsecase(todoRepo domain.Todo) DeleteTodoUsecase {
	return &deleteTodoUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (c *deleteTodoUsecaseImpl) Exec(ctx context.Context, id int) error {
	err := c.todoRepo.DeleteTodo(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
