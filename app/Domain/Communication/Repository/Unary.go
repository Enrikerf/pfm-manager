package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type Unary interface {
	Communicate(task Task.Task) (Content.Content, error)
}
