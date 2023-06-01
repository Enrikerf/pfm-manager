package Step

import "errors"

type InvalidSentenceLengthError interface {
	error
}

func NewInvalidSentenceLengthError() InvalidSentenceLengthError {
	return errors.New("step.sentence length must be less than 255")
}
