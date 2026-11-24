package sections

import (
	env "go-wallet/internal/shared/src/infrastructure/utils"
)

type LoggingConfig struct {
	Level  string
	Format string
}

func LoadLoggingConfig() (LoggingConfig, error) {
	level, err := env.ParseOrDefault(
		"LOG_LEVEL",
		"info",
		env.String,
	)
	if err != nil {
		return LoggingConfig{}, err
	}

	format, err := env.ParseOrDefault(
		"LOG_FORMAT",
		"json",
		env.String,
	)
	if err != nil {
		return LoggingConfig{}, err
	}

	return LoggingConfig{
		Level:  level,
		Format: format,
	}, nil
}
