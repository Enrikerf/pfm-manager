package ExecutionModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	CommunicationMode string
	ResultOK          ExecutionMode.Mode
	ResultKO          ExecutionMode.UnknownError
}{
	{
		CommunicationMode: "UNARY",
		ResultOK:          ExecutionMode.Automatic,
		ResultKO:          nil,
	},
	{
		CommunicationMode: "NOT_VALID",
		ResultKO:          ExecutionMode.NewUnknownError(),
	},
}

func TestFromString(t *testing.T) {
	for _, test := range tests {
		mode, err := ExecutionMode.FromString(test.CommunicationMode)
		if test.ResultKO != nil {
			assert.ErrorIs(t, err, test.ResultKO)
		} else {
			assert.Equal(t, test.ResultOK, mode)
		}

	}
}
