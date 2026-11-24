package sections

import (
	env "go-wallet/internal/shared/src/infrastructure/utils"
)

type ServerConfig struct {
	Host string
	Port int
}

func LoadServerConfig() (ServerConfig, error) {
	host, err := env.ParseOrDefault(
		"APP_HOST",
		"0.0.0.0",
		env.String,
	)
	if err != nil {
		return ServerConfig{}, err
	}

	port, err := env.ParseOrDefault(
		"APP_PORT",
		8080,
		env.Int,
	)
	if err != nil {
		return ServerConfig{}, err
	}

	return ServerConfig{
		Host: host,
		Port: port,
	}, nil
}
