package Error

import "errors"

type TaskNotFoundError interface {
	error
}

func NewTaskNotFoundError() TaskNotFoundError {
	return errors.New("task not found")
}
