package storage_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	awsClient "github.com/keito-isurugi/todo-app-backend/internal/infra/aws"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/env"
	"github.com/keito-isurugi/todo-app-backend/internal/infra/storage"
)

func Test_s3Repository_PutObject(t *testing.T) {
	tests := []struct {
		id      int
		name    string
		request *entity.PutObjectInput
		wantErr bool
	}{
		{
			id:      1,
			name:    "正常系",
			request: entity.NewPutObjectInput("test-key", "img/png", []byte("test-content")),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			ev, err := env.NewValue()
			a.NoError(err)

			s3Client, err := awsClient.NewAWSSession(ev)
			a.NoError(initTestBucket(s3Client, ev.AwsS3BucketName))

			s3Repo := storage.NewS3Repository(ev, s3Client)
			if err = s3Repo.PutObject(tt.request); (err != nil) != tt.wantErr {
				t.Errorf("s3Repository.PutObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPreSignedObject(t *testing.T) {
	tests := []struct {
		id      int
		name    string
		key     string
		wantErr bool
	}{
		{
			id:      1,
			name:    "正常系",
			key:     "test-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			a.NoError(os.Setenv("ENV", "test"))

			ev, err := env.NewValue()
			a.NoError(err)

			s3Client, err := awsClient.NewAWSSession(ev)
			a.NoError(initTestBucket(s3Client, ev.AwsS3BucketName))

			s3Repo := storage.NewS3Repository(ev, s3Client)
			url := s3Repo.GetPreSignedObject(tt.key)
			a.NotEmpty(url)
		})
	}
}

// initTestBucket
// テスト用のバケットを初期化する
// バケットが存在する場合はバケット内のオブジェクトを全て削除して、バケットを削除してから作成する
func initTestBucket(s3Client s3iface.S3API, bucket string) error {
	if existTestBucket(s3Client, bucket) {
		err := deleteTestBucket(s3Client, bucket)
		if err != nil {
			return err
		}
	}
	err := createTestBucket(s3Client, bucket)
	if err != nil {
		return err
	}

	return nil
}

// existTestBucket バケットが存在するかどうかを確認する
func existTestBucket(s3Client s3iface.S3API, bucket string) bool {
	_, err := s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return false
	}
	return true
}

// createTestBucket バケットを作成する
func createTestBucket(s3Client s3iface.S3API, bucket string) error {
	_, err := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeBucketAlreadyOwnedByYou {
			return nil
		}
		return err
	}
	return nil
}

// deleteTestBucket バケットの中のオブジェクトを全て削除してからバケットを削除する
func deleteTestBucket(s3Client s3iface.S3API, bucket string) error {
	listObjectsOutput, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	for _, object := range listObjectsOutput.Contents {
		deleteObjectInput := &s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		}
		_, err = s3Client.DeleteObject(deleteObjectInput)
		if err != nil {
			return err
		}
	}

	_, err = s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}
