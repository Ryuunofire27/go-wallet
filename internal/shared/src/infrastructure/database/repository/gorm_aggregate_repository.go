package repository

import (
	"context"
	"fmt"
	"go-wallet/internal/shared/src/application/port/events"
	"go-wallet/internal/shared/src/application/port/unit_of_work"
	"go-wallet/internal/shared/src/domain/aggregate"
	apperror "go-wallet/internal/shared/src/domain/error"
	"go-wallet/internal/shared/src/domain/vo"
	"strings"

	"gorm.io/gorm"
)

type GormAggregateRepository[T aggregate.Aggregate[K], K vo.ID, M any] struct {
	uow       unit_of_work.DBProvider[*gorm.DB]
	mapper    EntityMapper[T, M]
	collector events.EventCollector
	name      string
}

type AggregateMapper[T any, M any] interface {
	ToDomain(model *M) (*T, error)
	ToModel(aggregate T) (*M, error)
}

func NewGormAggregateRepository[T aggregate.Aggregate[K], K vo.ID, M any](
	uow unit_of_work.DBProvider[*gorm.DB],
	mapper EntityMapper[T, M],
	name string,
) *GormAggregateRepository[T, K, M] {
	return &GormAggregateRepository[T, K, M]{
		uow:    uow,
		mapper: mapper,
		name:   name,
	}
}

func (r *GormAggregateRepository[T, K, M]) DB() *gorm.DB {
	return r.uow.DB()
}

func (r *GormAggregateRepository[T, K, M]) Create(ctx context.Context, aggregate *T) (*T, error) {
	model, err := r.mapper.ToModel(aggregate)

	if err != nil {
		return nil, err
	}

	result := r.DB().WithContext(ctx).Create(model)

	if result.Error != nil {
		return nil, apperror.Wrap(
			apperror.Internal,
			"ENTITY_NOT_SAVED",
			fmt.Sprintf("Entity from %s not saved", strings.ToUpper(r.name)),
			fmt.Sprintf("%s.Save", r.name),
			result.Error,
		)
	}

	r.collector.Add((*aggregate).Collect()...)

	return r.mapper.ToDomain(model)
}
