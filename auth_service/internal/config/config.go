package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/call_center_request/auth_service/pkg/logging"
	"time"
)

type Config struct {
	Grpc     Grpc                 `env:"GRPC"`
	Logging  logging.LoggerConfig `env:"LOG"`
	Postgres PostgreSQL           `env:"POSTGRES"`
}

type Grpc struct {
	Host string `env:"GRPC_HOST" required:"true" default:"0.0.0.0"`
	Port string `env:"GRPC_PORT" required:"true" default:"50000"`
}

type PostgreSQL struct {
	PostgresqlHost     string        `env:"POSTGRES_HOST" default:"localhost"`
	PostgresqlPort     string        `env:"POSTGRES_PORT" default:"5432"`
	PostgresqlUser     string        `env:"POSTGRES_USER" default:"postgres"`
	PostgresqlPassword string        `env:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresqlDBName   string        `env:"POSTGRES_DBNAME" default:"postgres"`
	MaxIdleConnTime    time.Duration `env:"POSTGRES_MAX_IDLE_CONN_TIME" default:"5m"`
	MaxConns           int           `env:"POSTGRES_MAX_CONNS" default:"20"`
	ConnMaxLifetime    time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME" default:"10m"`
}

func New() *Config {
	cfg := Config{}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(fmt.Errorf("[config.Env] read env: %w", err))
	}

	return &cfg
}
