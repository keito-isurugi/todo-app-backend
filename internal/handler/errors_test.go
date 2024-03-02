package handler

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	dbError "github.com/keito-isurugi/todo-app-backend/internal/infra/postgres"
)

func TestCreateErrResponse(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		name     string
		err      error
		expected errResponse
	}{
		{
			name:     "NotFoundError",
			err:      &dbError.NotFoundError{Message: "not found error"},
			expected: errResponse{Message: "not found error", Status: http.StatusNotFound},
		},
		{
			name:     "BadRequestError",
			err:      &dbError.BadRequestError{Message: "bad request error"},
			expected: errResponse{Message: "bad request error", Status: http.StatusBadRequest},
		},
		{
			name:     "DuplicateError",
			err:      &dbError.DuplicateError{Message: "duplicate error"},
			expected: errResponse{Message: "duplicate error", Status: http.StatusConflict},
		},
		{
			name:     "DuplicateUniqueKeyError",
			err:      &dbError.DuplicateUniqueKeyError{Message: "duplicate unique key error"},
			expected: errResponse{Message: "duplicate unique key error", Status: http.StatusConflict},
		},
		{
			name:     "OtherError",
			err:      fmt.Errorf("other error"),
			expected: errResponse{Message: "other error", Status: http.StatusInternalServerError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := createErrResponse(tt.err)
			a.Equal(tt.expected, actual)
		})
	}
}
