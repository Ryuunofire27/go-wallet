package database

import "time"

const (
	MAX_OPEN_CONNECTIONS = 10
	MAX_IDLE_CONNECTIONS = 5
	MAX_LIFE_TIME        = 30 * time.Minute
	MAX_IDLE_TIME        = 5 * time.Minute
)
