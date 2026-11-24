package sections

import (
	env "go-wallet/internal/shared/src/infrastructure/utils"
	"time"
)

type SecurityConfig struct {
	JWTSecret     string
	JWTExpiration time.Duration
}

func LoadSecurityConfig() (SecurityConfig, error) {
	jwtSecret, err := env.Parse(
		"JWT_SECRET",
		env.String,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	jwtExpiration, err := env.ParseOrDefault(
		"JWT_EXPIRATION",
		15*time.Minute,
		env.Duration,
	)
	if err != nil {
		return SecurityConfig{}, err
	}

	return SecurityConfig{
		JWTSecret:     jwtSecret,
		JWTExpiration: jwtExpiration,
	}, nil
}
