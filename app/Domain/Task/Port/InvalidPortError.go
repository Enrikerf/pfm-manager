package Port

import "errors"

type InvalidPortError interface {
	error
}

func NewInvalidPortError() InvalidPortError {
	return errors.New("invalid host format")
}
