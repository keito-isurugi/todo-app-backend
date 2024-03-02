package usecase

import (
	"context"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
)

type ChangeTodoDoneFlagUsecase interface {
	Exec(ctx context.Context, dh *entity.Todo) error
}

type changeTodoDoneFlagUsecaseImpl struct {
	todoRepo domain.Todo
}

func NewChangeTodoDoneFlagUsecase(todoRepo domain.Todo) ChangeTodoDoneFlagUsecase {
	return &changeTodoDoneFlagUsecaseImpl{
		todoRepo: todoRepo,
	}
}

func (c *changeTodoDoneFlagUsecaseImpl) Exec(ctx context.Context, todo *entity.Todo) error {
	err := c.todoRepo.ChangeTodoDoneFlag(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}
