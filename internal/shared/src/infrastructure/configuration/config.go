package config

import "go-wallet/internal/shared/src/infrastructure/configuration/sections"

type Config struct {
	Environment string

	Server   sections.ServerConfig
	Database sections.DatabaseConfig
	Security sections.SecurityConfig
	Logging  sections.LoggingConfig
}
