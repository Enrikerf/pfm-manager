package TaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/CommunicationModeTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/ExecutionModeTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/HostTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/PortTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/StepTest"
)

type TaskBuilder interface {
	WithId(id Task.Id) TaskBuilder
	WithStatus(status Status.Status) TaskBuilder
	Build() Task.Task
}

type taskBuilder struct {
	id                Task.Id
	host              Host.Vo
	port              Port.Vo
	steps             []Step.Vo
	communicationMode CommunicationMode.Mode
	executionMode     ExecutionMode.Mode
	status            Status.Status
}

func Get() TaskBuilder {
	return &taskBuilder{
		id:                DefaultId(),
		host:              HostTest.BuildDefaultMock(),
		port:              PortTest.BuildDefaultMock(),
		steps:             []Step.Vo{StepTest.BuildDefaultStepVo()},
		communicationMode: CommunicationModeTest.BuildDefaultMock(),
		executionMode:     ExecutionModeTest.BuildDefaultMock(),
		status:            Status.New(Status.Pending),
	}
}

func (taskBuilder *taskBuilder) WithId(id Task.Id) TaskBuilder {
	taskBuilder.id = id
	return taskBuilder
}

func (taskBuilder *taskBuilder) WithStatus(status Status.Status) TaskBuilder {
	taskBuilder.status = status
	return taskBuilder
}

func (taskBuilder *taskBuilder) Build() Task.Task {
	task, _ := Task.Load(
		taskBuilder.id,
		taskBuilder.host,
		taskBuilder.port,
		taskBuilder.steps,
		taskBuilder.communicationMode,
		taskBuilder.executionMode,
		taskBuilder.status,
	)
	return task
}

func BuildDefaultTaskMock() Task.Task {
	task, _ := Task.New(
		HostTest.BuildDefaultMock(),
		PortTest.BuildDefaultMock(),
		[]Step.Vo{StepTest.BuildDefaultStepVo()},
		CommunicationModeTest.BuildDefaultMock(),
		ExecutionModeTest.BuildDefaultMock(),
	)
	return task
}
