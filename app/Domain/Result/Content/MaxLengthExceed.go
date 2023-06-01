package Content

import "errors"

type MaxLengthExceed interface {
	error
}

func NewMaxLengthExceed() MaxLengthExceed {
	return errors.New("content max length exceed")
}
