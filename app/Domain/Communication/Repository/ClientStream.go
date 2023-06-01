package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type ClientStream interface {
	Communicate(task Task.Task) (Content.Content, error)
}
