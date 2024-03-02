package main

import (
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/keito-isurugi/todo-app-backend/internal/infra/aws"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/logger"
	"github.com/keito-isurugi/todo-app-backend/internal/server"
)

//	@Summary		Swagger Example API
//	@version		v1
//	@description	BA Portal Replace API
//	@host			localhost

func main() {
	ev, _ := env.NewValue()
	zapLogger, _ := logger.NewLogger(true)
	defer func() { _ = zapLogger.Sync() }()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              ev.Sentry.Dsn,
		TracesSampleRate: ev.TracesSampleRate,
		Environment:      ev.Env, // 環境名
		AttachStacktrace: true,   // CaptureMessageを呼んだ時にスタックトレースを付与するかのオプション
	})
	if err != nil {
		zapLogger.Error(err.Error())
	}

	defer func() {
		sentry.Recover()
		sentry.Flush(2 * time.Second)
	}()

	dbClient, err := db.NewClient(&ev.DB, zapLogger)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	awsClient, err := aws.NewAWSSession(ev)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	router := server.SetupRouter(ev, dbClient, awsClient, zapLogger)

	router.Logger.Fatal(router.Start(":" + ev.ServerPort))
}
