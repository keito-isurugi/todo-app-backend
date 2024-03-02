package db

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
)

type Client interface {
	Conn(ctx context.Context) *gorm.DB
}

type client struct {
	db *gorm.DB
}

func NewClient(e *env.DB, zapLogger *zap.Logger) (Client, error) {
	gormLogger := initGormLogger(zapLogger)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		e.PostgresHost,
		e.PostgresUser,
		e.PostgresPassword,
		e.PostgresDatabase,
		e.PostgresPort,
	)
	log.Printf("dsn: %s", dsn)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}
	db.Logger = db.Logger.LogMode(gormLogger.LogLevel)

	if err != nil {
		return nil, err
	}
	return &client{
		db: db,
	}, nil
}

func (c *client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}
