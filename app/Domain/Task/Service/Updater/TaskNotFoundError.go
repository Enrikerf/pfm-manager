package Updater

import "errors"

type TaskNotFoundError interface {
	error
}

func NewTaskNotFoundError() TaskNotFoundError {
	return errors.New("NewTaskNotFoundError")
}
