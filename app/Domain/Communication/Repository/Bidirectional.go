package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Bidirectional interface {
	Setup(host Host.Vo, port Port.Vo) error
	Write(step Step.Step) error
	Read() (Content.Content, error)
	Close() error
}
