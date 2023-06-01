package UpdateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Updater"
)

type UseCase interface {
	Update(command Command) error
}

func New(
	findRepository Repository.Find,
	saveRepository Repository.Save,
	dispatcher Event.Dispatcher,
) UseCase {
	return &useCase{Updater.Updater{
		FindRepository: findRepository,
		SaveRepository: saveRepository,
		Dispatcher:     dispatcher,
	}}
}

type useCase struct {
	updater Updater.Updater
}

func (service *useCase) Update(command Command) error {
	id, err := Task.LoadIdFromString(command.Uuid)
	if err != nil {
		return err
	}
	err = service.updater.Update(id, command.Host, command.Port, command.Status)
	if err != nil {
		return err
	}
	return nil
}
