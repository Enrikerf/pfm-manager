package ExecutionModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	hostString string
	ResultKO   Host.InvalidHostError
}{
	{
		hostString: "127.12.12.1",
		ResultKO:   nil,
	},
	{
		hostString: "NOT_VALID",
		ResultKO:   Host.NewInvalidHostError(),
	},
}

func TestFromString(t *testing.T) {
	for _, test := range tests {
		host, err := Host.NewVo(test.hostString)
		if test.ResultKO != nil {
			assert.ErrorIs(t, err, test.ResultKO)
		} else {
			assert.Equal(t, test.hostString, host.GetValue())
		}
	}
}
