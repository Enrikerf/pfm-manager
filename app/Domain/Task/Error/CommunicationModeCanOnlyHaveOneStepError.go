package Error

import "errors"

type CommunicationModeCanOnlyHaveOneStepError interface {
	error
}

func NewCommunicationModeCanOnlyHaveOneStepError() CommunicationModeCanOnlyHaveOneStepError {
	return errors.New("NewCommunicationModeCanOnlyHaveOneStepError")
}
