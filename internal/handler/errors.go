package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	dbError "github.com/keito-isurugi/todo-app-backend/internal/infra/postgres"
)

type errResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func newErrResponse(code int, message string) *errResponse {
	return &errResponse{
		Message: message,
		Status:  code,
	}
}

func createErrResponse(err error) errResponse {
	switch err := err.(type) {
	case *dbError.NotFoundError:
		return errResponse{
			Message: err.Message,
			Status:  http.StatusNotFound,
		}
	case *dbError.BadRequestError:
		return errResponse{
			Message: err.Message,
			Status:  http.StatusBadRequest,
		}
	case *dbError.DuplicateError:
		return errResponse{
			Message: err.Message,
			Status:  http.StatusConflict,
		}
	case *dbError.DuplicateUniqueKeyError:
		return errResponse{
			Message: err.Message,
			Status:  http.StatusConflict,
		}
	case validator.ValidationErrors:
		return errResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}
	return errResponse{Message: err.Error(), Status: http.StatusInternalServerError}
}

func (r *errResponse) outputErrorLog(zapLogger *zap.Logger, message, traceID string) {
	msg := fmt.Sprintf("%s: %s", message, r.Message)
	zapLogger.Error(msg, zap.String("trace_id", traceID))
}
