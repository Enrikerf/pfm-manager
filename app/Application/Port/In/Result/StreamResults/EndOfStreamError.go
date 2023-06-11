package StreamResults

import "errors"

type EndOfStreamError interface {
	error
}

func NewEndOfStreamError() EndOfStreamError {
	return errors.New("EndOfStreamError")
}
