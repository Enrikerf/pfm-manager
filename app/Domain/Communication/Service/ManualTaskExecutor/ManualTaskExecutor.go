package ManualTaskExecutor

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/BidirectionalCommunicator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/ClientStreamCommunicator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/ServerStreamCommunicator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Service/UnaryCommunicator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Error"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
)

type Execute interface {
	Execute(taskId Task.Id) (Result.Batch, error)
}

func New(
	finder Finder.Finder,
	saveTaskRepository TaskRepository.Save,
	saveBatchRepository ResultRepository.SaveBatch,
	bidirectionalCommunicator BidirectionalCommunicator.BidirectionalCommunicator,
	unaryCommunicator UnaryCommunicator.UnaryCommunicator,
	serverStreamCommunicator ServerStreamCommunicator.ServerStreamCommunicator,
	clientStreamCommunicator ClientStreamCommunicator.ClientStreamCommunicator,
) Execute {
	return &executor{
		finder,
		saveTaskRepository,
		saveBatchRepository,
		bidirectionalCommunicator,
		unaryCommunicator,
		serverStreamCommunicator,
		clientStreamCommunicator,
	}
}

type executor struct {
	taskFinder                Finder.Finder
	saveTaskRepository        TaskRepository.Save
	saveBatchRepository       ResultRepository.SaveBatch
	bidirectionalCommunicator BidirectionalCommunicator.BidirectionalCommunicator
	unaryCommunicator         UnaryCommunicator.UnaryCommunicator
	serverStreamCommunicator  ServerStreamCommunicator.ServerStreamCommunicator
	clientStreamCommunicator  ClientStreamCommunicator.ClientStreamCommunicator
}

func (e *executor) Execute(taskId Task.Id) (Result.Batch, error) {
	task, err := e.taskFinder.Find(taskId)
	if err != nil {
		return nil, err
	}
	if task.GetExecutionMode() != ExecutionMode.Manual {
		return nil, Error.NewTaskNotManualCanNotBeExecutedManuallyError()
	}
	if task.GetStatus().Value() == Status.Running {
		return nil, Error.NewRunningTaskCanNotBeExecutedManuallyError()
	}
	batch := Result.NewBatch(task.GetId())
	e.saveBatchRepository.Persist(batch)
	task.SetStatus(Status.New(Status.Running))
	e.saveTaskRepository.Persist(task)
	switch task.GetCommunicationMode().Value() {
	case CommunicationMode.Unary:
		go e.unaryCommunicator.Communicate(task, batch)
	case CommunicationMode.Bidirectional:
		go e.bidirectionalCommunicator.Communicate(task, batch)
	case CommunicationMode.ServerStream:
		go e.serverStreamCommunicator.Communicate(task, batch)
	case CommunicationMode.ClientStream:
		go e.clientStreamCommunicator.Communicate(task, batch)
	}
	return batch, nil
}
