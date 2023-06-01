package UpdateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type Command struct {
	Uuid              string
	Host              Host.Vo
	Port              Port.Vo
	CommunicationMode CommunicationMode.Mode
	ExecutionMode     ExecutionMode.Mode
	Status            Status.Status
}
