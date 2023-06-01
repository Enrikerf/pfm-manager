package Error

import "errors"

type BatchNotFoundError interface {
	error
}

func NewBatchNotFoundError() BatchNotFoundError {
	return errors.New("batch not found")
}
