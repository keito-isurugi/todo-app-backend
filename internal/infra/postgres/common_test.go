package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
)

func truncateTable(ctx context.Context, t *testing.T, dbClient db.Client) {
	a := assert.New(t)

	// Start a new transaction
	tx := dbClient.Conn(ctx).Begin()

	// Execute TRUNCATE statements within the transaction
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.banks RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.branches RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.menu_masters RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.menu_groups RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.menu_master_group RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.appointments RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.global_holidays RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.local_holidays RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.global_local_holiday RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.daily_hours RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.auto_mail_settings RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.weekly_hours RESTART IDENTITY CASCADE").Error)
	a.NoError(tx.Exec("TRUNCATE TABLE todo_app_test.public.banner_images RESTART IDENTITY CASCADE").Error)

	// Commit the transaction
	a.NoError(tx.Commit().Error)
}

type mockDBClient struct {
	mockDB *gorm.DB
}

func (m *mockDBClient) Conn(context.Context) *gorm.DB {
	return m.mockDB
}

func NewMockDBClient(db *gorm.DB) db.Client {
	return &mockDBClient{mockDB: db}
}
