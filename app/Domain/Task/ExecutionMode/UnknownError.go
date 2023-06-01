package ExecutionMode

import "errors"

type UnknownError interface {
	error
}

func NewUnknownError() UnknownError {
	return errors.New("execution mode not valid")
}
