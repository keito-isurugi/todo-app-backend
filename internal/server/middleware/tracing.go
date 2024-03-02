package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewTracing() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			traceID, err := generateTraceID()
			if err != nil {
				traceID = "unknown"
			}
			ctx.Set("trace_id", traceID)
			return next(ctx)
		}
	}
}

func generateTraceID() (string, error) {
	traceID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return traceID.String(), nil
}

func getTraceID(ctx echo.Context) string {
	traceID := ctx.Get("trace_id")
	if traceID == nil {
		traceID, err := generateTraceID()
		if err != nil {
			return ""
		}
		return traceID
	}
	return traceID.(string)
}
