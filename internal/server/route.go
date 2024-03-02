package server

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/keito-isurugi/todo-app-backend/internal/handler"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/db"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/postgres"
	"github.com/keito-isurugi/todo-app-backend/internal/server/middleware"
)

func SetupRouter(ev *env.Values, dbClient db.Client, awsClient s3iface.S3API, zapLogger *zap.Logger) *echo.Echo {
	e := echo.New()
	// middleware
	e.Use(middleware.NewLogging(zapLogger))
	e.Use(middleware.NewTracing())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	//postgres
	todoRepo := postgres.NewTodoRepository(dbClient, zapLogger)

	//handler
	todoHandler := handler.NewTodoHandler(todoRepo, zapLogger)

	// todos
	todoGroup := e.Group("/todos")
	todoGroup.GET("", todoHandler.ListTodos)
	todoGroup.POST("/:id", todoHandler.RegisterTodo)
	todoGroup.PUT("/:id/_change", todoHandler.ChangeTodo)

	for _, route := range e.Routes() {
		zapLogger.Info("route", zap.String("method", route.Method), zap.String("path", route.Path))
	}

	return e
}
