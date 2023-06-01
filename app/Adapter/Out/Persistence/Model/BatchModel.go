package Model

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/google/uuid"
	"time"
)

type BatchDb struct {
	ID        uint
	Uuid      uuid.UUID
	TaskID    uint
	TaskUuid  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BatchDb) TableName() string {
	return "batches"
}

func (batchModel *BatchDb) FromDomainV2(batch Result.Batch) {
	batchModel.Uuid = batch.GetId().GetUuid()
	batchModel.TaskUuid = batch.GetTaskId().GetUuid()
	batchModel.CreatedAt = batch.GetCreatedAt()
}

func (batchModel *BatchDb) ToDomainV2() (Result.Batch, error) {
	var results []Result.Result
	selfEntity := Result.LoadBatch(
		Result.LoadBatchId(batchModel.Uuid),
		Task.LoadId(batchModel.TaskUuid),
		results,
		batchModel.CreatedAt,
	)
	return selfEntity, nil
}
