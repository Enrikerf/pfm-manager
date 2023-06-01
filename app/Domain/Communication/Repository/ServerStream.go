package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
)

type ServerStream interface {
	Setup(host Host.Vo, port Port.Vo) error
	Communicate(task Task.Task) error
	GetIterator() (Content.Content, error)
}
