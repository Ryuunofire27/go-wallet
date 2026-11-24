package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Parser[T any] func(string) (T, error)

func Parse[T any](key string, parser Parser[T]) (T, error) {
	value, exists := os.LookupEnv(key)

	var zero T

	if !exists || value == "" {
		return zero, fmt.Errorf("%s: environment variable is required", key)
	}

	parsed, err := parser(value)
	if err != nil {
		return zero, fmt.Errorf("%s: invalid value %q: %w", key, value, err)
	}

	return parsed, nil
}

func ParseOrDefault[T any](
	key string,
	defaultValue T,
	parser Parser[T],
) (T, error) {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		return defaultValue, nil
	}

	parsed, err := parser(value)
	if err != nil {
		var zero T
		return zero, fmt.Errorf(
			"%s: invalid value %q: %w",
			key,
			value,
			err,
		)
	}

	return parsed, nil
}

func String(value string) (string, error) {
	return value, nil
}

func Int(value string) (int, error) {
	return strconv.Atoi(value)
}

func Bool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func Duration(value string) (time.Duration, error) {
	return time.ParseDuration(value)
}
