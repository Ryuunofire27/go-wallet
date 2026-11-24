package repository

import (
	"context"
	"fmt"
	"go-wallet/internal/shared/src/application/port/unit_of_work"
	"go-wallet/internal/shared/src/domain/entity"
	apperror "go-wallet/internal/shared/src/domain/error"
	"go-wallet/internal/shared/src/domain/vo"
	"strings"

	"gorm.io/gorm"
)

type GormEntityRepository[T entity.DomainEntity[K], K vo.ID, M any] struct {
	uow    unit_of_work.DBProvider[*gorm.DB]
	mapper EntityMapper[T, M]
	name   string
}

type EntityMapper[T any, M any] interface {
	ToDomain(model *M) (*T, error)
	ToModel(entity *T) (*M, error)
}

func NewGormEntityRepository[T entity.DomainEntity[K], K vo.ID, M any](
	uow unit_of_work.DBProvider[*gorm.DB],
	mapper EntityMapper[T, M],
	name string,
) *GormEntityRepository[T, K, M] {
	return &GormEntityRepository[T, K, M]{
		uow:    uow,
		mapper: mapper,
		name:   name,
	}
}

func (r *GormEntityRepository[T, K, M]) DB() *gorm.DB {
	return r.uow.DB()
}

func (r *GormEntityRepository[T, K, M]) Create(ctx context.Context, entity *T) (*T, error) {
	model, err := r.mapper.ToModel(entity)

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

	return r.mapper.ToDomain(model)
}
