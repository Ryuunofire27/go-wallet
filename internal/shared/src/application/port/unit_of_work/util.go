package unit_of_work

import "context"

func Transaction[T any](
	ctx context.Context,
	uow TransactionManager,
	fn func(ctx context.Context) (T, error),
) (T, error) {
	var result T

	err := uow.Transaction(ctx, func(ctx context.Context) error {
		var callbackErr error
		result, callbackErr = fn(ctx)
		return callbackErr
	})

	if err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}
