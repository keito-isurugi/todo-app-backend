package postgres

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
)

type todoRepository struct {
	dbClient  db.Client
	zapLogger *zap.Logger
}

func NewTodoRepository(dbClient db.Client, zapLogger *zap.Logger) domain.Todo {
	return &todoRepository{
		dbClient:  dbClient,
		zapLogger: zapLogger,
	}
}

func (r *todoRepository) ListTodos(ctx context.Context) (entity.ListTodos, error) {
	var todos entity.ListTodos
	if err := r.dbClient.Conn(ctx).
		Table("todos").
		Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) GetTodo(ctx context.Context, id int) (*entity.Todo, error) {
	var todo entity.Todo
	if err := r.dbClient.Conn(ctx).
		Table("todos").
		Where("id", id).
		First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &NotFoundError{Message: err.Error()}
		}
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepository) RegisterTodo(ctx context.Context, todo *entity.Todo) (int, error) {
	if err := t.dbClient.Conn(ctx).Table("todos").Create(todo).Error; err != nil {
		return 0, err
	}
	return todo.ID, nil
}

func (t *todoRepository) ChangeTodoDoneFlag(ctx context.Context, todo *entity.Todo) error {
	var to *entity.Todo
	if err := t.dbClient.Conn(ctx).
		Table("todos").
		Where("id", todo.ID).
		First(&to).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &NotFoundError{Message: err.Error()}
		}
		return err
	}

	updateColumns := map[string]any{
		"done_flag": todo.DoneFlag,
	}

	return t.dbClient.Conn(ctx).
		Model(&entity.Todo{}).
		Where("id = ?", todo.ID).
		Updates(updateColumns).Error
}

func (t *todoRepository) DeleteTodo(ctx context.Context, id int) error {
	var todo *entity.Todo
	if err := t.dbClient.Conn(ctx).
		Table("todos").
		Where("id", id).
		First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &NotFoundError{Message: err.Error()}
		}
		return err
	}

	if err := t.dbClient.Conn(ctx).Delete(&todo).Error; err != nil {
        return err
    }

    return nil
}