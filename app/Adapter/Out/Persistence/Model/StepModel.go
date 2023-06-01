package Model

import (
	StepDomain "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"time"
)

type StepDb struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	Sentence  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StepDb) TableName() string {
	return "steps"
}

func (stepModel *StepDb) FromDomainV2(taskUuid uuid.UUID, selfEntity StepDomain.Step) {
	stepModel.Uuid = selfEntity.GetId().GetUuid()
	stepModel.TaskUuid = taskUuid
	stepModel.Sentence = selfEntity.GetSentence()
}

func (stepModel *StepDb) ToDomainV2() (StepDomain.Step, error) {

	vo, err := StepDomain.NewVo(stepModel.Sentence)
	if err != nil {
		return nil, err
	}
	selfEntity := StepDomain.Load(
		StepDomain.LoadId(stepModel.Uuid),
		vo,
	)

	return selfEntity, nil
}
