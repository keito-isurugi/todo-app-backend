package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/infra/postgres"
)

func TestNotFoundError_Error(t *testing.T) {
	e := &postgres.NotFoundError{
		Message: "Not Found",
	}
	assert.Equal(t, "Not Found", e.Error())
}

func TestBadRequestError_Error(t *testing.T) {
	e := &postgres.BadRequestError{
		Message: "Bad Request",
	}
	assert.Equal(t, "Bad Request", e.Error())
}

func TestDuplicateError_Error(t *testing.T) {
	e := &postgres.DuplicateError{
		Message: "Duplicate",
	}
	assert.Equal(t, "Duplicate", e.Error())
}

func TestDuplicateUniqueKeyError_Error(t *testing.T) {
	e := &postgres.DuplicateUniqueKeyError{
		Message: "Duplicate Unique Key",
	}
	assert.Equal(t, "Duplicate Unique Key", e.Error())
}

func TestInternalServerError_Error(t *testing.T) {
	e := &postgres.InternalServerError{
		Message: "Internal Server Error",
	}
	assert.Equal(t, "Internal Server Error", e.Error())
}

func TestCheckConstraint_Error(t *testing.T) {
	e := &postgres.CheckConstraint{
		Message: "Check Constraint",
	}
	assert.Equal(t, "Check Constraint", e.Error())
}
