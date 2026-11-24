package entity

import "go-wallet/internal/shared/src/domain/vo"

type Entity[T vo.ID] struct {
	ID T
}

func (e Entity[T]) GetID() T {
	return e.ID
}

type DomainEntity[K vo.ID] interface {
	GetID() K
}
