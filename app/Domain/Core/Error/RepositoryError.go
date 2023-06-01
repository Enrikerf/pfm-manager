package Error

import "errors"

type RepositoryError interface {
	error
}

func NewRepositoryError(msg string) RepositoryError {
	return errors.New(msg)
}
