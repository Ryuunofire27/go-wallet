package events

import "time"

type Metadata struct {
	ID               string
	Name             string
	OccurredAt       time.Time
	AggregateID      string
	AggregateName    string
	AggregateVersion uint64
}

type DomainEvent[T any] struct {
	Metadata Metadata
	Payload  T
}

func (e DomainEvent[T]) EventMetadata() Metadata {
	return e.Metadata
}

type Event interface {
	EventMetadata() Metadata
}
