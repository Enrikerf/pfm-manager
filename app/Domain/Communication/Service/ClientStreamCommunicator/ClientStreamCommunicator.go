package ClientStreamCommunicator

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
)

type ClientStreamCommunicator interface {
	Communicate(task Task.Task, batch Result.Batch)
}

func New(
	clientStreamRepository Repository.ClientStream,
	saveResultRepository ResultRepository.Save,
) ClientStreamCommunicator {
	return &serverStreamCommunicator{
		clientStreamRepository: clientStreamRepository,
		saveResultPort:         saveResultRepository,
	}
}

type serverStreamCommunicator struct {
	clientStreamRepository Repository.ClientStream
	finder                 Finder.Finder
	saveResultPort         ResultRepository.Save
}

func (ssc *serverStreamCommunicator) Communicate(task Task.Task, batch Result.Batch) {
	content, err := ssc.clientStreamRepository.Communicate(task)
	if err != nil {
		ssc.processError(err, batch)
		return
	}
	result := Result.New(batch.GetId(), content)
	ssc.saveResultPort.Persist(result)

}

func (ssc *serverStreamCommunicator) processError(err error, batch Result.Batch) {
	content, err := Content.NewContent(err.Error())
	if err != nil {
		content, err = Content.NewContent(err.Error())
	}
	result := Result.New(batch.GetId(), content)
	ssc.saveResultPort.Persist(result)
}
