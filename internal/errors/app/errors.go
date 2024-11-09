package apperrors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrAlreadyExists      = errors.New("already exists")
	ErrPreconditionFailed = errors.New("precondition failed")
	ErrUnauthenticated    = errors.New("unauthenticated")
)

func NewErrNotFound(err string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, err)
}

func NewErrInvalidArgument(err string) error {
	return fmt.Errorf("%w: %s", ErrInvalidArgument, err)
}

func NewErrAlreadyExists(err string) error {
	return fmt.Errorf("%w: %s", ErrAlreadyExists, err)
}

func NewErrPreconditionFailed(err string) error {
	return fmt.Errorf("%w: %s", ErrPreconditionFailed, err)
}

func NewErrUnauthenticated(err string) error {
	return fmt.Errorf("%w: %s", ErrUnauthenticated, err)
}
