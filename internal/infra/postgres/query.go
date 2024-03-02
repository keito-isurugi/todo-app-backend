package postgres

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/keito-isurugi/todo-app-backend/internal/helper"
)

func addWhereEq(query *gorm.DB, columnName string, value any) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s = ?", columnName), value)
}

func addWhereGte(query *gorm.DB, columnName string, value any) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s >= ?", columnName), value)
}

func addWhereLte(query *gorm.DB, columnName string, value any) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s <= ?", columnName), value)
}

func addWhereLike(query *gorm.DB, columnName string, value string) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", columnName), "%"+value+"%")
}
