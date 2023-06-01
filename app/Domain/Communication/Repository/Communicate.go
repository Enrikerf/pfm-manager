package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type Communicate interface {
	Communicate(task Task.Task) Result.Batch
}
