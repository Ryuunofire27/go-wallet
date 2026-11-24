package apperror

import "errors"

type Kind string

const (
	NotFound     Kind = "NOT_FOUND"
	Conflict     Kind = "CONFLICT"
	Validation   Kind = "VALIDATION_ERROR"
	Unauthorized Kind = "UNAUTHORIZED"
	Forbidden    Kind = "FORBIDDEN"
	Internal     Kind = "INTERNAL_ERROR"
)

type AppError struct {
	Kind      Kind
	Code      string
	Message   string
	Operation string
	Cause     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) Is(target error) bool {
	targetErr, ok := target.(*AppError)
	if !ok {
		return false
	}

	return e.Kind == targetErr.Kind &&
		e.Code == targetErr.Code
}

func New(kind Kind, code, message string, operation string) *AppError {
	return &AppError{
		Kind:      kind,
		Code:      code,
		Message:   message,
		Operation: operation,
	}
}

func Wrap(
	kind Kind,
	code string,
	message string,
	operation string,
	cause error,
) *AppError {
	return &AppError{
		Kind:      kind,
		Code:      code,
		Message:   message,
		Operation: operation,
		Cause:     cause,
	}
}

func IsKind(err error, kind Kind) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Kind == kind
}
