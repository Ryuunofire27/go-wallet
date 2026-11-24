package database

import (
	"time"

	"gorm.io/gorm"
)

type PoolConfig struct {
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifetime        time.Duration
	MaxIdleTime        time.Duration
}

type ConnectionsConfig struct {
	ReadDSN  string
	WriteDSN string

	ReadPool  PoolConfig
	WritePool PoolConfig
}

type Connections struct {
	Read  *gorm.DB
	Write *gorm.DB
}
