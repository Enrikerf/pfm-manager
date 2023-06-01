package DeleteTask

import (
	"errors"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Eraser"
	"github.com/google/uuid"
)

type UseCase interface {
	Delete(command Command) error
}

func New(findRepository Repository.Find, deleteRepository Repository.Delete) UseCase {
	return &useCase{Eraser.Eraser{FindRepository: findRepository, DeleteRepository: deleteRepository}}
}

type useCase struct {
	eraser Eraser.Eraser
}

func (service *useCase) Delete(command Command) error {
	uuidToDelete, err := uuid.Parse(command.Uuid)
	if err != nil {
		return errors.New("WRONG ID")
	}
	err = service.eraser.Erase(Task.LoadId(uuidToDelete))
	if err != nil {
		return err
	}
	return nil
}
