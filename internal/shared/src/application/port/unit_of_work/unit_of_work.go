package unit_of_work

import "context"

type TransactionManager interface {
	Transaction(
		ctx context.Context,
		fn func(ctx context.Context) error,
	) error
}

type DBProvider[T any] interface {
	DB() T
}

type UnitOfWork[T any] interface {
	TransactionManager
	DBProvider[T]
}
