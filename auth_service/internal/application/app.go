package application

import (
	"database/sql"
	"fmt"
	"github.com/lallison21/call_center_request/auth_service/internal/config"
	"github.com/lallison21/call_center_request/auth_service/internal/repository"
	"github.com/lallison21/call_center_request/auth_service/internal/service"
	"github.com/lallison21/call_center_request/auth_service/pkg/logging"
	"github.com/lallison21/call_center_request/auth_service/version"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Application struct {
	log         *zerolog.Logger
	authService *service.Service
}

func New(cfg *config.Config) *Application {
	log := logging.New(cfg.Logging)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlDBName,
	)

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())

	authRepo := repository.New(db)
	authService := service.New(authRepo)

	return &Application{log: log, authService: authService}
}

func (a *Application) Run() {
	a.log.Info().Msgf("service name: %s; service version: %s; service commit: %s; service build time: %s",
		version.Name,
		version.Version,
		version.Commit,
		version.BuildTime,
	)
}
