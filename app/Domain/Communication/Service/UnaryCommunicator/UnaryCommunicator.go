package UnaryCommunicator

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
)

type UnaryCommunicator interface {
	Communicate(task Task.Task, batch Result.Batch)
}

func New(
	unaryRepository Repository.Unary,
	saveResultRepository ResultRepository.Save,
) UnaryCommunicator {
	return &unaryCommunicator{
		unaryRepository: unaryRepository,
		saveResultPort:  saveResultRepository,
	}
}

type unaryCommunicator struct {
	unaryRepository Repository.Unary
	finder          Finder.Finder
	saveResultPort  ResultRepository.Save
}

func (unaryCommunicator *unaryCommunicator) Communicate(task Task.Task, batch Result.Batch) {
	content, err := unaryCommunicator.unaryRepository.Communicate(task)
	if err != nil {
		unaryCommunicator.processError(err, batch)
		return
	}
	result := Result.New(batch.GetId(), content)
	unaryCommunicator.saveResultPort.Persist(result)
}

func (unaryCommunicator *unaryCommunicator) processError(err error, batch Result.Batch) {
	content, err := Content.NewContent(err.Error())
	if err != nil {
		content, err = Content.NewContent(err.Error())
	}
	result := Result.New(batch.GetId(), content)
	unaryCommunicator.saveResultPort.Persist(result)
}
