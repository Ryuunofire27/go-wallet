package unit_of_work

import (
	"context"
	"errors"

	uowApp "go-wallet/internal/shared/src/application/port/unit_of_work"
	apperror "go-wallet/internal/shared/src/domain/error"

	"gorm.io/gorm"
)

type GormUnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewGormUnitOfWork(db *gorm.DB) *GormUnitOfWork {
	return &GormUnitOfWork{
		db: db,
	}
}

func (uow *GormUnitOfWork) DB() *gorm.DB {
	if uow.tx != nil {
		return uow.tx
	}

	return uow.db
}

func (uow *GormUnitOfWork) Transaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	if uow.tx != nil {
		return apperror.New(
			apperror.Internal,
			uowApp.TransactionAlreadyActive,
			"Transaction already active",
			"GormUnitOfWork.Transaction",
		)
	}

	tx := uow.db.WithContext(ctx).Begin()

	if tx.Error != nil {
		return apperror.Wrap(
			apperror.Internal,
			uowApp.TransactionBeginFailed,
			"Could not begin transaction",
			"GormUnitOfWork.Transaction",
			tx.Error,
		)
	}

	uow.tx = tx

	defer func() {
		uow.tx = nil
	}()

	err := fn(ctx)

	if err != nil {
		rollbackErr := uow.tx.Rollback().Error

		if rollbackErr != nil {
			return apperror.Wrap(
				apperror.Internal,
				uowApp.TransactionRollbackFailed,
				"Could not rollback transaction",
				"GormUnitOfWork.Transaction",
				errors.Join(err, rollbackErr),
			)
		}

		return err
	}

	commitErr := uow.tx.Commit().Error

	if commitErr != nil {
		return apperror.Wrap(
			apperror.Internal,
			uowApp.TransactionCommitFailed,
			"Could not commit transaction",
			"GormUnitOfWork.Transaction",
			commitErr,
		)
	}

	return nil
}
