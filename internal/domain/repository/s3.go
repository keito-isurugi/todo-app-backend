//go:generate mockgen -source=s3.go -destination=./mock/s3_mock.go
package domain

import "github.com/keito-isurugi/todo-app-backend/internal/domain/entity"

type S3Repository interface {
	PutObject(input *entity.PutObjectInput) error
	GetPreSignedObject(key string) string
}
