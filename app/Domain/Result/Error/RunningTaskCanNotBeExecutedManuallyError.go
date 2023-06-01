package Error

import "errors"

type RunningTaskCanNotBeExecutedManuallyError interface {
	error
}

func NewRunningTaskCanNotBeExecutedManuallyError() RunningTaskCanNotBeExecutedManuallyError {
	return errors.New("RunningTaskCanNotBeExecutedManuallyError")
}
