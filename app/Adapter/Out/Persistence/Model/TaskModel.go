package Model

import (
	TaskDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"time"
)

type TaskDb struct {
	ID            uint
	Uuid          uuid.UUID
	Host          string
	Port          string
	Steps         []StepDb `gorm:"foreignKey:TaskID;references:ID"`
	Mode          string
	Status        string
	ExecutionMode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (TaskDb) TableName() string {
	return "tasks"
}

func (taskDb *TaskDb) FromDomainV2(selfEntity TaskDomain.Task) {
	taskDb.Uuid = selfEntity.GetId().GetUuid()
	taskDb.Host = selfEntity.GetHost().GetValue()
	taskDb.Port = selfEntity.GetPort().GetValue()
	for _, stepDomain := range selfEntity.GetSteps() {
		step := StepDb{}
		step.FromDomainV2(taskDb.Uuid, stepDomain)
		taskDb.Steps = append(taskDb.Steps, step)
	}
	taskDb.Mode = string(selfEntity.GetCommunicationMode())
	taskDb.Status = string(selfEntity.GetStatus().Value())
	taskDb.ExecutionMode = string(selfEntity.GetExecutionMode())
}

func (taskDb *TaskDb) ToDomainV2() (TaskDomain.Task, error) {
	host, err := Host.NewVo(taskDb.Host)
	if err != nil {
		return nil, err
	}
	port, err := Port.NewVo(taskDb.Port)
	if err != nil {
		return nil, err
	}
	communicationMode, err := CommunicationMode.FromString(taskDb.Mode)
	if err != nil {
		return nil, err
	}
	executionMode, err := ExecutionMode.FromString(taskDb.ExecutionMode)
	if err != nil {
		return nil, err
	}
	status, err := Status.FromString(taskDb.Status)
	if err != nil {
		return nil, err
	}
	var stepVos []Step.Vo
	for _, commandSentence := range taskDb.Steps {
		stepVo, err := Step.NewVo(commandSentence.Sentence)
		if err != nil {
			return nil, err
		}
		stepVos = append(stepVos, stepVo)
	}
	if err != nil {
		return nil, err
	}
	task, err := TaskDomain.Load(
		TaskDomain.LoadId(taskDb.Uuid),
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
		status,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}
