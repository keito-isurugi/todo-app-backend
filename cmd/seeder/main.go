package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/logger"
)

type Todo struct {
	Title string
}

func main() {
	ev, _ := env.NewValue()

	zapLogger, _ := logger.NewLogger(true)
	defer func() { _ = zapLogger.Sync() }()

	dbClient, err := db.NewClient(&ev.DB, zapLogger)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	gofakeit.Seed(time.Now().UnixNano())

	var todos []Todo // Change the way you create the slice.

	for i := 0; i < 50; i++ {
		todo := Todo{
			Title: "タイトル" + fmt.Sprint(i),
		}
		todos = append(todos, todo) // Add each appointment to the slice.
	}

	dbClient.Conn(context.Background()).Exec("TRUNCATE todos RESTART IDENTITY CASCADE")
	if err := dbClient.Conn(context.Background()).Create(&todos).Error; err != nil {
		fmt.Println(err)
	}
}
