package BidirectionalCommunicator

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type BidirectionalCommunicator interface {
	Communicate(task Task.Task, batch Result.Batch)
}

func New(
	find TaskRepository.Find,
	callBidiPort Repository.Bidirectional,
	saveResultRepository ResultRepository.Save,
) BidirectionalCommunicator {
	return &bidirectionalCommunicator{
		finder:         Finder.Finder{FindRepository: find},
		callBidiPort:   callBidiPort,
		saveResultPort: saveResultRepository,
	}
}

type bidirectionalCommunicator struct {
	finder         Finder.Finder
	callBidiPort   Repository.Bidirectional
	saveResultPort ResultRepository.Save
}

func (b *bidirectionalCommunicator) Communicate(task Task.Task, batch Result.Batch) {

	err := b.callBidiPort.Setup(task.GetHost(), task.GetPort())
	if err != nil {
		b.processError(err, batch)
		return
	}
	if !b.write(task.GetSteps()[0], batch) {
		b.close(batch)
		return
	}
	for {
		task, _ := b.finder.Find(batch.GetTaskId())
		if task.GetStatus().Value() != Status.Running {
			if len(task.GetSteps()) > 1 && !b.write(task.GetSteps()[1], batch) {
				b.close(batch)
				return
			}
			break
		}
		resultsContent, err := b.read()
		if err != nil {
			b.processError(err, batch)
			break
		}
		if resultsContent != nil {
			result := Result.New(batch.GetId(), resultsContent)
			b.saveResultPort.Persist(result)
		}
	}
	b.close(batch)
}

func (b *bidirectionalCommunicator) write(step Step.Step, batch Result.Batch) bool {
	err := b.callBidiPort.Write(step)
	if err != nil {
		b.processError(err, batch)
		return false
	}
	return true
}

func (b *bidirectionalCommunicator) read() (Content.Content, error) {
	resultsContent, err := b.callBidiPort.Read()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//if len(resultsContent.GetValue()) > 0 {
	//	return nil, nil
	//}
	return resultsContent, nil
}

func (b *bidirectionalCommunicator) close(batch Result.Batch) {
	err := b.callBidiPort.Close()
	if err != nil {
		b.processError(err, batch)
	}
}

func (b *bidirectionalCommunicator) processError(err error, batch Result.Batch) {
	content, err := Content.NewContent(err.Error())
	if err != nil {
		fmt.Println(err.Error())
		content, err = Content.NewContent(err.Error())
	}
	result := Result.New(batch.GetId(), content)
	b.saveResultPort.Persist(result)
}
