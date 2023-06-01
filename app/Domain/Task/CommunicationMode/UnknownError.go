package CommunicationMode

import "errors"

type UnknownError interface {
	error
}

func NewUnknownError() UnknownError {
	return errors.New("communication mode not valid")
}
