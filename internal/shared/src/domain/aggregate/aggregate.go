package aggregate

import (
	"go-wallet/internal/shared/src/domain/entity"
	"go-wallet/internal/shared/src/domain/events"
	"go-wallet/internal/shared/src/domain/vo"
	"time"
)

type AggregateRoot[T vo.ID] struct {
	entity.Entity[T]

	domainEvents     []events.Event
	AggregateVersion uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func (ar *AggregateRoot[T]) AddDomainEvent(event events.Event) {
	ar.domainEvents = append(ar.domainEvents, event)
}

func (ar *AggregateRoot[T]) Collect() []events.Event {
	collected := ar.domainEvents
	ar.domainEvents = nil

	return collected
}

type Aggregate[K vo.ID] interface {
	entity.DomainEntity[K]
	Collect() []events.Event
}
