package Looper

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Communication/Repository"
	ResultRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	TaskRepository "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"sync"
)

var once sync.Once
var instance Looper

type Looper interface {
	IsEnabled() bool
	Enable()
}

func NewLooper(
	communicateRepository Repository.Communicate,
	findTasksByRepository TaskRepository.FindBy,
	saveTaskRepository TaskRepository.Save,
	saveBatchRepository ResultRepository.SaveBatch,
	saveResultRepository ResultRepository.Save,
) Looper {
	once.Do(func() {
		instance = &looper{
			communicateRepository: communicateRepository,
			findTasksByRepository: findTasksByRepository,
			saveTaskRepository:    saveTaskRepository,
			saveBatchRepository:   saveBatchRepository,
			saveResultRepository:  saveResultRepository,
			isLoopEnabled:         make(chan bool, 1),
		}
	})
	return instance
}

type looper struct {
	communicateRepository Repository.Communicate
	findTasksByRepository TaskRepository.FindBy
	saveTaskRepository    TaskRepository.Save
	saveBatchRepository   ResultRepository.SaveBatch
	saveResultRepository  ResultRepository.Save
	isLoopEnabled         chan bool
}

func (l *looper) IsEnabled() bool {
	return len(l.isLoopEnabled) != 0
}

func (l *looper) Enable() {
	l.isLoopEnabled <- true
	go l.loop()
}

func (l *looper) loop() {
	for l.IsEnabled() {
		tasks, err := l.findTasksByRepository.FindBy(map[string]interface{}{
			"status":         Status.Pending,
			"execution_mode": ExecutionMode.Automatic,
		})
		if err != nil {
			fmt.Printf(err.Error())
			l.stopLoop()
			return
		}
		if len(tasks) < 1 {
			l.stopLoop()
			return
		}

		var wg sync.WaitGroup
		for index := range tasks {
			wg.Add(1)
			go l.executeTask(&wg, tasks[index])
		}
		wg.Wait()
	}
}

func (l *looper) executeTask(wg *sync.WaitGroup, task Task.Task) {
	defer wg.Done()
	task.SetStatus(Status.New(Status.Running))
	l.saveTaskRepository.Persist(task)
	resultBatch := l.communicateRepository.Communicate(task)
	l.saveBatchRepository.Persist(resultBatch)
	task.SetStatus(Status.New(Status.Done))
	l.saveTaskRepository.Persist(task)
}

func (l *looper) stopLoop() {
	if len(l.isLoopEnabled) > 0 {
		fmt.Println("loop stopped")
		<-l.isLoopEnabled
	} else {
		fmt.Println("trying to stop loop but is not running")
	}
}
