package ServerStreamCommunicator

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
)

type ServerStreamCommunicator interface {
	Communicate(task Task.Task, batch Result.Batch)
}

func New(
	serverStreamRepository Repository.ServerStream,
	saveResultRepository ResultRepository.Save,
) ServerStreamCommunicator {
	return &serverStreamCommunicator{
		serverStreamRepository: serverStreamRepository,
		saveResultPort:         saveResultRepository,
	}
}

type serverStreamCommunicator struct {
	serverStreamRepository Repository.ServerStream
	finder                 Finder.Finder
	saveResultPort         ResultRepository.Save
}

func (ssc *serverStreamCommunicator) Communicate(task Task.Task, batch Result.Batch) {
	err := ssc.serverStreamRepository.Setup(task.GetHost(), task.GetPort())
	if err != nil {
		ssc.processError(err, batch)
		return
	}
	err = ssc.serverStreamRepository.Communicate(task)
	if err != nil {
		ssc.processError(err, batch)
		return
	}
	for {
		content, err := ssc.serverStreamRepository.GetIterator()
		if err == nil && content == nil {
			return
		}
		if err != nil {
			ssc.processError(err, batch)
			return
		}
		result := Result.New(batch.GetId(), content)
		ssc.saveResultPort.Persist(result)
	}

}

func (ssc *serverStreamCommunicator) processError(err error, batch Result.Batch) {
	content, err := Content.NewContent(err.Error())
	if err != nil {
		content, err = Content.NewContent(err.Error())
	}
	result := Result.New(batch.GetId(), content)
	ssc.saveResultPort.Persist(result)
}
