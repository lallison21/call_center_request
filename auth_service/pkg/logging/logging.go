package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type LoggerConfig struct {
	LogIndex  string `env:"LOG_INDEX" required:"true" default:"my_app"`
	IsDebug   bool   `env:"LOG_IS_DEBUG" required:"true" default:"false"`
	LogToFile bool   `env:"LOG_TO_FILE" required:"true" default:"true"`
}

func New(cfg LoggerConfig) *zerolog.Logger {
	logger := log.With().Str("service", cfg.LogIndex).Logger()

	if cfg.IsDebug {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	if cfg.LogToFile {
		// В файл
		file, err := os.Create("application.log")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to create log file")
		}
		logger = logger.Output(file)
	} else {
		// В os.STDOUT
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return &logger
}
