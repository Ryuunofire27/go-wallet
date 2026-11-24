package database

import "fmt"

func NewConnections(cfg ConnectionsConfig) (*Connections, error) {

	writeDB, err := NewConnection(
		cfg.WriteDSN,
		cfg.WritePool,
	)
	if err != nil {
		return nil, fmt.Errorf("connect write database: %w", err)
	}

	readDB, err := NewConnection(
		cfg.ReadDSN,
		cfg.ReadPool,
	)
	if err != nil {
		sqlDB, _ := writeDB.DB()
		_ = sqlDB.Close()

		return nil, fmt.Errorf("connect read database: %w", err)
	}

	return &Connections{
		Write: writeDB,
		Read:  readDB,
	}, nil

}
