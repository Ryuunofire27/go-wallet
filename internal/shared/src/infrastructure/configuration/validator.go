package config

import (
	"fmt"
)

func Validate(cfg *Config) error {
	if cfg.Environment == "" {
		return fmt.Errorf("APP_ENV is required")
	}

	if err := cfg.Database.WritePool.Validate(); err != nil {
		return fmt.Errorf("write pool: %w", err)
	}

	if err := cfg.Database.ReadPool.Validate(); err != nil {
		return fmt.Errorf("read pool: %w", err)
	}

	if cfg.Server.Port < 1 || cfg.Server.Port > 65535 {
		return fmt.Errorf("server port must be between 1 and 65535")
	}

	if cfg.Security.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	return nil
}
