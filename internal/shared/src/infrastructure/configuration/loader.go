package config

import (
	"fmt"
	"go-wallet/internal/shared/src/infrastructure/configuration/sections"
	env "go-wallet/internal/shared/src/infrastructure/utils"
)

func Load() (*Config, error) {
	cfg := &Config{}

	environment, err := env.ParseOrDefault(
		"APP_ENV",
		"development",
		env.String,
	)
	if err != nil {
		return nil, err
	}

	cfg.Environment = environment

	databaseConfig, err := sections.LoadDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("database config: %w", err)
	}

	cfg.Database = databaseConfig

	serverConfig, err := sections.LoadServerConfig()

	if err != nil {
		return nil, fmt.Errorf("server config: %w", err)
	}

	cfg.Server = serverConfig

	securityConfig, err := sections.LoadSecurityConfig()
	if err != nil {
		return nil, fmt.Errorf("server config: %w", err)
	}

	cfg.Security = securityConfig

	loggingConfig, err := sections.LoadLoggingConfig()

	if err != nil {
		return nil, fmt.Errorf("logging config: %w", err)
	}

	cfg.Logging = loggingConfig

	if err := Validate(cfg); err != nil {
		return nil, fmt.Errorf("configuration validation: %w", err)
	}

	return cfg, nil
}
