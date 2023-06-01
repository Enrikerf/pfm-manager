package Repository

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"

type Save interface {
	Persist(task Task.Task)
}
