package Repository

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"

//go:generate mockery
type Delete interface {
	Delete(id Task.Id) error
}
