package database

import (
	config "go-wallet/internal/shared/src/infrastructure/configuration"
)

func Connect(config config.Config) (*Connections, error) {
	return NewConnections(ConnectionsConfig{
		WriteDSN:  config.Database.WriteDSN,
		ReadDSN:   config.Database.ReadDSN,
		WritePool: PoolConfig(config.Database.WritePool),
		ReadPool:  PoolConfig(config.Database.ReadPool),
	})
}
