package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
)

type Task interface {
	GetId() Id
	GetHost() Host.Vo
	GetPort() Port.Vo
	GetSteps() []Step.Step
	GetCommunicationMode() CommunicationMode.Mode
	GetExecutionMode() ExecutionMode.Mode
	GetStatus() Status.Status
	SetHost(host Host.Vo)
	SetPort(port Port.Vo)
	SetStatus(status Status.Status)
}

type task struct {
	id                Id
	host              Host.Vo
	port              Port.Vo
	steps             []Step.Step
	communicationMode CommunicationMode.Mode
	executionMode     ExecutionMode.Mode
	status            Status.Status
}

func New(
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) (Task, error) {
	err := validateInputs(stepVos, communicationMode, executionMode)
	if err != nil {
		return nil, err
	}
	task := &task{}
	task.id = NewId()
	for _, stepVo := range stepVos {
		step := Step.New(stepVo)
		task.steps = append(task.steps, step)
	}
	task.host = host
	task.port = port
	task.executionMode = executionMode
	task.communicationMode = communicationMode
	task.status = Status.New(Status.Pending)
	return task, nil
}
func validateInputs(
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode) error {
	if len(stepVos) < 1 {
		return Error.NewTaskMustHaveAtLeastOneStepError()
	}
	if (communicationMode.Value() == CommunicationMode.Unary || communicationMode.Value() == CommunicationMode.ServerStream) &&
		len(stepVos) > 1 {
		return Error.NewCommunicationModeCanOnlyHaveOneStepError()
	}
	if executionMode == ExecutionMode.Manual && communicationMode.Value() == CommunicationMode.Bidirectional &&
		len(stepVos) > 2 {
		return Error.NewManualBidirectionalTaskOnlyCanHave2StepsError()
	}
	return nil
}

func Load(
	id Id,
	host Host.Vo,
	port Port.Vo,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
	status Status.Status,
) (Task, error) {
	task := &task{}
	err := validateInputs(stepVos, communicationMode, executionMode)
	if err != nil {
		return nil, err
	}
	task.id = id
	for _, stepVo := range stepVos {
		step := Step.New(stepVo)
		task.steps = append(task.steps, step)
	}
	task.host = host
	task.port = port
	task.executionMode = executionMode
	task.communicationMode = communicationMode
	task.status = status
	return task, nil
}

func (t *task) GetId() Id {
	return t.id
}

func (t *task) GetHost() Host.Vo {
	return t.host
}

func (t *task) SetHost(host Host.Vo) {
	t.host = host
}

func (t *task) GetPort() Port.Vo {
	return t.port
}

func (t *task) SetPort(port Port.Vo) {
	t.port = port
}

func (t *task) GetSteps() []Step.Step {
	return t.steps
}

func (t *task) GetCommunicationMode() CommunicationMode.Mode {
	return t.communicationMode
}

func (t *task) GetExecutionMode() ExecutionMode.Mode {
	return t.executionMode
}

func (t *task) GetStatus() Status.Status {
	return t.status
}

func (t *task) SetStatus(status Status.Status) {
	t.status = status
}
