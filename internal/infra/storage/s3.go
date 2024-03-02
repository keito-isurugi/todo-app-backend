package storage

import (
	"bytes"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
	"github.com/keito-isurugi/todo-app-backend/internal/helper"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
)

type s3Repository struct {
	ev *env.Values
	s3 s3iface.S3API
}

func NewS3Repository(ev *env.Values, s3 s3iface.S3API) domain.S3Repository {
	return &s3Repository{
		ev: ev,
		s3: s3,
	}
}

func (r *s3Repository) PutObject(poi *entity.PutObjectInput) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(r.ev.AwsS3BucketName),
		Key:         aws.String(poi.Key),
		ContentType: aws.String(poi.ContentType),
		Body:        bytes.NewReader(poi.FileContent),
	}
	_, err := r.s3.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}

func (r *s3Repository) GetPreSignedObject(key string) string {
	req, _ := r.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(r.ev.AwsS3BucketName),
		Key:    aws.String(key),
	})

	urlStr, _ := req.Presign(15 * time.Minute)
	// ローカル環境の場合、署名を削除
	if helper.IsLocal() {
		strings.Split(urlStr, "?")
		url := strings.Split(urlStr, "?")[0]
		return strings.Replace(url, "localstack", "localhost", 1)
	}
	return urlStr
}
