package postgres_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/logger"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/postgres"
)

func TestNewTodoRepository_ListTodos(t *testing.T) {
	tests := []struct {
		id     int
		name   string
		todoID string
		want   entity.ListTodos
		setup  func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:     1,
			name:   "正常系/一覧取得",
			todoID: "101",
			want: entity.ListTodos{
				{
					ID:        1,
					Title:     "タイトル1",
					DoneFlag:  false,
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					// DeletedAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Title:     "タイトル2",
					DoneFlag:  true,
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					// DeletedAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO todo_app_test.public.todos (id,title,created_at, updated_at, deleted_at) VALUES ('1','タイトル1','false','2021-01-01T00:00:00+00:00','admin', '2021-01-01T00:00:00+00:00','null')`).Error)
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO todo_app_test.public.todos (id,title,created_at, updated_at, deleted_at) VALUES ('2','タイトル2','true','2021-01-01T00:00:00+00:00','admin', '2021-01-01T00:00:00+00:00','null')`).Error)
			},
		},
		{
			id:   2,
			name: "対象のレコードがない場合、空の配列が返ってくること",
			want: entity.ListTodos{},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			zapLogger, err := logger.NewLogger(true)
			a.NoError(err)

			ev, err := env.NewValue()
			a.NoError(err)

			dbClient, err := db.NewClient(&ev.DB, zapLogger)
			a.NoError(err)

			truncateTable(ctx, t, dbClient)

			if tt.setup != nil {
				tt.setup(ctx, t, dbClient)
			}

			todoRepo := postgres.NewTodoRepository(dbClient, zapLogger)

			got, err := todoRepo.ListTodos(ctx)
			a.NoError(err)

			if tt.want != nil {
				if !cmp.Equal(got, tt.want) {
					t.Errorf("diff %s", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

func TestNewTodoRepository_RegisterTodo(t *testing.T) {
	tests := []struct {
		id        int
		name      string
		request   *entity.Todo
		wantTable *entity.Todo
		wantError error
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/新規作成",
			request: &entity.Todo{
				Title: "タイトル1",
			},
			wantTable: &entity.Todo{
				ID:        1,
				Title:     "タイトル1",
				DoneFlag:  false,
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				// UpdatedAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO todo_app_test.public.todos (id,title,created_at, updated_at, deleted_at) VALUES ('1','タイトル1','false','2021-01-01T00:00:00+00:00','admin', '2021-01-01T00:00:00+00:00','null')`).Error)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			zapLogger, err := logger.NewLogger(true)
			a.NoError(err)

			ev, err := env.NewValue()
			a.NoError(err)

			dbClient, err := db.NewClient(&ev.DB, zapLogger)
			a.NoError(err)

			truncateTable(ctx, t, dbClient)

			if tt.setup != nil {
				tt.setup(ctx, t, dbClient)
			}

			todoRepo := postgres.NewTodoRepository(dbClient, zapLogger)
			_, err = todoRepo.RegisterTodo(ctx, tt.request)
			a.NoError(err)

			var got *entity.Todo
			if err = dbClient.Conn(ctx).Find(&got).Error; err != nil {
				a.EqualError(err, tt.wantError.Error())
				return
			}

			if tt.wantTable != nil {
				// ID, CreatedAt, UpdatedAtはランダムの値が入ってくるので比較対象外
				opt := cmpopts.IgnoreFields(entity.Todo{}, "ID", "CreatedAt", "UpdatedAt")
				if !cmp.Equal(got, tt.wantTable, opt) {
					t.Errorf("diff %s", cmp.Diff(got, tt.wantTable))
				}
			}
		})
	}
}

func TestNewTodoRepository_ChangeTodo(t *testing.T) {
	tests := []struct {
		id        int
		name      string
		request   *entity.Todo
		wantTable *entity.Todo
		wantError error
		setup     func(ctx context.Context, t *testing.T, dbClient db.Client)
	}{
		{
			id:   1,
			name: "正常系/更新",
			request: &entity.Todo{
				ID:    1,
				Title: "タイトル1",
			},
			wantTable: &entity.Todo{
				ID:        1,
				Title:     "タイトル1",
				DoneFlag:  false,
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			setup: func(ctx context.Context, t *testing.T, dbClient db.Client) {
				assert.NoError(t, dbClient.Conn(ctx).Exec(`INSERT INTO todo_app_test.public.todos (id,title,created_at, updated_at, deleted_at) VALUES ('1','タイトル1','false','2021-01-01T00:00:00+00:00','admin', '2021-01-01T00:00:00+00:00','null')`).Error)
			},
		},
		{
			id:   2,
			name: "異常系/レコードが見つからない場合、NotFoundエラーになること",
			request: &entity.Todo{
				ID: 2,
			},
			wantTable: nil,
			wantError: &postgres.NotFoundError{Message: "record not found"},
			setup:     func(ctx context.Context, t *testing.T, dbClient db.Client) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			zapLogger, err := logger.NewLogger(true)
			a.NoError(err)

			ev, err := env.NewValue()
			a.NoError(err)

			dbClient, err := db.NewClient(&ev.DB, zapLogger)
			a.NoError(err)

			truncateTable(ctx, t, dbClient)

			if tt.setup != nil {
				tt.setup(ctx, t, dbClient)
			}

			todoRepo := postgres.NewTodoRepository(dbClient, zapLogger)
			err = todoRepo.ChangeTodo(ctx, tt.request)

			if tt.wantError != nil {
				a.EqualError(err, tt.wantError.Error())
				return
			}

			a.NoError(err)
			var got *entity.Todo
			if err = dbClient.Conn(ctx).Find(&got).Error; err != nil {
				a.EqualError(err, tt.wantError.Error())
				return
			}

			if tt.wantTable != nil {
				// UpdatedAtはランダムの値が入ってくるので比較対象外
				opt := cmpopts.IgnoreFields(entity.Todo{}, "UpdatedAt")
				if !cmp.Equal(got, tt.wantTable, opt) {
					t.Errorf("diff %s", cmp.Diff(got, tt.wantTable))
				}
			}
		})
	}
}
