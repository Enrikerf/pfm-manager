package ExecutionModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
)

func BuildDefaultMock() ExecutionMode.Mode {
	return ExecutionMode.New(ExecutionMode.Manual)
}
