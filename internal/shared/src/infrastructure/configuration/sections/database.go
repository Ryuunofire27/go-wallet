package sections

import (
	"errors"
	env "go-wallet/internal/shared/src/infrastructure/utils"
	"time"
)

type PoolConfig struct {
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifetime        time.Duration
	MaxIdleTime        time.Duration
}

type DatabaseConfig struct {
	WriteDSN string
	ReadDSN  string

	WritePool PoolConfig
	ReadPool  PoolConfig
}

func (p PoolConfig) Validate() error {
	if p.MaxOpenConnections <= 0 {
		return errors.New("max open connections must be greater than zero")
	}

	if p.MaxIdleConnections < 0 {
		return errors.New("max idle connections cannot be negative")
	}

	if p.MaxIdleConnections > p.MaxOpenConnections {
		return errors.New("max idle connections cannot exceed max open connections")
	}

	if p.MaxLifetime < 0 || p.MaxIdleTime < 0 {
		return errors.New("pool durations cannot be negative")
	}

	return nil
}

func LoadDatabaseConfig() (DatabaseConfig, error) {
	writeDSN, err := env.Parse("WRITE_DSN", env.String)
	if err != nil {
		return DatabaseConfig{}, err
	}

	readDSN, err := env.Parse("READ_DSN", env.String)
	if err != nil {
		return DatabaseConfig{}, err
	}

	writePool, err := loadPoolConfig("WRITE")
	if err != nil {
		return DatabaseConfig{}, err
	}

	readPool, err := loadPoolConfig("READ")
	if err != nil {
		return DatabaseConfig{}, err
	}

	return DatabaseConfig{
		WriteDSN:  writeDSN,
		ReadDSN:   readDSN,
		WritePool: writePool,
		ReadPool:  readPool,
	}, nil
}

func loadPoolConfig(prefix string) (PoolConfig, error) {
	maxOpen, err := env.ParseOrDefault(
		prefix+"_MAX_OPEN_CONNECTIONS",
		10,
		env.Int,
	)
	if err != nil {
		return PoolConfig{}, err
	}

	maxIdle, err := env.ParseOrDefault(
		prefix+"_MAX_IDLE_CONNECTIONS",
		5,
		env.Int,
	)
	if err != nil {
		return PoolConfig{}, err
	}

	maxLifetime, err := env.ParseOrDefault(
		prefix+"_MAX_LIFETIME",
		30*time.Minute,
		env.Duration,
	)
	if err != nil {
		return PoolConfig{}, err
	}

	maxIdleTime, err := env.ParseOrDefault(
		prefix+"_MAX_IDLE_TIME",
		5*time.Minute,
		env.Duration,
	)
	if err != nil {
		return PoolConfig{}, err
	}

	return PoolConfig{
		MaxOpenConnections: maxOpen,
		MaxIdleConnections: maxIdle,
		MaxLifetime:        maxLifetime,
		MaxIdleTime:        maxIdleTime,
	}, nil
}
