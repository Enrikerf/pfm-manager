package Error

import "errors"

type TaskMustHaveAtLeastOneStepError interface {
	error
}

func NewTaskMustHaveAtLeastOneStepError() TaskMustHaveAtLeastOneStepError {
	return errors.New("NewTaskMustHaveAtLeastOneStepError")
}
