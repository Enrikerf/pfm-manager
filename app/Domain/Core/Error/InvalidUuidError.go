package Error

import "errors"

type InvalidUuidError interface {
	error
}

func NewInvalidUuidError() InvalidUuidError {
	return errors.New("not parseable ID")
}
