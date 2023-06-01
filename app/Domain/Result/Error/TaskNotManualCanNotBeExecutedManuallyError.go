package Error

import "errors"

type TaskNotManualCanNotBeExecutedManuallyError interface {
	error
}

func NewTaskNotManualCanNotBeExecutedManuallyError() TaskNotManualCanNotBeExecutedManuallyError {
	return errors.New("NewTaskNotManualCanNotBeExecutedManuallyError")
}
