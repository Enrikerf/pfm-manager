package CommunicationModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	sentence string
	ResultKO error
}{
	{
		sentence: "sentence",
		ResultKO: nil,
	},
	{
		sentence: "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
		ResultKO: Step.NewInvalidSentenceLengthError(),
	},
}

func TestFromString(t *testing.T) {
	for _, test := range tests {
		vo, err := Step.NewVo(test.sentence)
		if test.ResultKO != nil {
			assert.Error(t, err)
			assert.Equal(t, test.ResultKO.Error(), err.Error())
		} else {
			assert.Equal(t, test.sentence, vo.GetSentence())
		}

	}
}
