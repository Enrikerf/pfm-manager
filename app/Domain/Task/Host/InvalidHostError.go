package Host

import "errors"

type InvalidHostError interface {
	error
}

func NewInvalidHostError() InvalidHostError {
	return errors.New("invalid host format")
}
