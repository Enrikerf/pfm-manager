package Eraser

import "errors"

type RunningTaskCantBeDeleted interface {
	error
}

func NewRunningTaskCantBeDeletedError() RunningTaskCantBeDeleted {
	return errors.New("NewRunningTaskCantBeDeletedError")
}
