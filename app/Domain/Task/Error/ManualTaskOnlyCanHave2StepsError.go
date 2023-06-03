package Error

import "errors"

type ManualBidirectionalTaskOnlyCanHave2StepsError interface {
	error
}

func NewManualBidirectionalTaskOnlyCanHave2StepsError() ManualBidirectionalTaskOnlyCanHave2StepsError {
	return errors.New("NewManualBidirectionalTaskOnlyCanHave2StepsError")
}
