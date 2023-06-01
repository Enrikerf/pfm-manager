package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/google/uuid"
	"time"
)

type ResultDb struct {
	ID        uint
	Uuid      uuid.UUID
	BatchID   uint
	BatchUuid uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ResultDb) TableName() string {
	return "results"
}

func (resultModel *ResultDb) FromDomainV2(result Result.Result) {
	resultModel.Uuid = result.GetId().GetUuid()
	resultModel.BatchUuid = result.GetBatchId().GetUuid()
	resultModel.Content = result.GetContent().GetValue()
	resultModel.CreatedAt = result.GetCreateAt()
}

func (resultModel *ResultDb) ToDomainV2() (Result.Result, error) {
	content, err := Content.NewContent(resultModel.Content)
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	selfEntity := Result.Load(
		Result.LoadId(resultModel.Uuid),
		Result.LoadBatchId(resultModel.BatchUuid),
		content,
		resultModel.CreatedAt,
	)
	return selfEntity, nil
}
