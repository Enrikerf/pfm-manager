package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type Find interface {
	Find(id Task.Id) (Task.Task, error)
}
