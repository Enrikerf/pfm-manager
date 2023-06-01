package CommunicationModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	CommunicationMode string
	ResultOK          CommunicationMode.Mode
	ResultKO          CommunicationMode.UnknownError
}{
	{
		CommunicationMode: "UNARY",
		ResultOK:          CommunicationMode.Unary,
		ResultKO:          nil,
	},
	{
		CommunicationMode: "NOT_VALID",
		ResultKO:          CommunicationMode.NewUnknownError(),
	},
}

func TestFromString(t *testing.T) {
	for _, test := range tests {
		communicationMode, err := CommunicationMode.FromString(test.CommunicationMode)
		if test.ResultKO != nil {
			assert.ErrorIs(t, err, test.ResultKO)
		} else {
			assert.Equal(t, test.ResultOK, communicationMode)
		}

	}
}
