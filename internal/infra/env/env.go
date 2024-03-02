package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Values struct {
	Env string `default:"local" split_words:"true"`
	Server
	DB
	TestDB
	AWS
	Sentry
	Debug bool `default:"true" split_words:"true"`
}

type Server struct {
	ServerPort string `default:"8080" split_words:"true"`
}

type DB struct {
	PostgresHost     string `required:"true" split_words:"true"`
	PostgresPort     string `required:"true" split_words:"true"`
	PostgresDatabase string `required:"true" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}

type TestDB struct {
	TestPostgresHost     string `required:"true" split_words:"true"`
	TestPostgresPort     string `required:"true" split_words:"true"`
	TestPostgresDatabase string `required:"true" split_words:"true"`
	TestPostgresUser     string `required:"true" split_words:"true"`
	TestPostgresPassword string `required:"true" split_words:"true"`
}

type AWS struct {
	AwsRegion          string `required:"true" split_words:"true"`
	AwsAccessKeyID     string `split_words:"true"`
	AwsSecretAccessKey string `split_words:"true"`
	AwsEndpoint        string `split_words:"true"`
	AwsEndpointLocal   string `split_words:"true"`
	AwsS3BucketName    string `required:"true" split_words:"true"`
}

type Sentry struct {
	Dsn              string  `default:"dummy" split_words:"true"`
	TracesSampleRate float64 `default:"1.0" split_words:"true"`
}

func NewValue() (*Values, error) {
	var v Values
	err := envconfig.Process("", &v)

	// test環境の場合はtest用のDB情報を設定する
	if v.Env == "test" {
		v.DB.PostgresHost = v.TestDB.TestPostgresHost
		v.DB.PostgresPort = v.TestDB.TestPostgresPort
		v.DB.PostgresDatabase = v.TestDB.TestPostgresDatabase
		v.DB.PostgresUser = v.TestDB.TestPostgresUser
		v.DB.PostgresPassword = v.TestDB.TestPostgresPassword
	}

	if err != nil {
		s := fmt.Sprintf("need to set all env values %+v", v)
		return nil, errors.Wrap(err, s)
	}
	return &v, nil
}
