package dependencies

import (
	"go-wallet/internal/shared/src/application/port/logger"
	config "go-wallet/internal/shared/src/infrastructure/configuration"
	"go-wallet/internal/shared/src/infrastructure/database"
)

type Dependencies struct {
	Connections *database.Connections
	Logger      logger.Logger
	Config      *config.Config
}

func NewDependencies(
	connections *database.Connections,
	logger logger.Logger,
	config *config.Config,
) *Dependencies {
	return &Dependencies{
		Connections: connections,
		Logger:      logger,
		Config:      config,
	}
}
