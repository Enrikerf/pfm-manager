package Repository

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
)

type FindBy interface {
	FindBy(conditions interface{}) ([]Task.Task, error)
}
