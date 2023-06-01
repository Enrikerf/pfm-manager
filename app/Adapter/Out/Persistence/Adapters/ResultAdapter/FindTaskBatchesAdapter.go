package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	Error2 "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"gorm.io/gorm"
)

type FindTaskBatchesAdapter struct {
	Orm *gorm.DB
}

func (adapter FindTaskBatchesAdapter) Find(id Task.Id) ([]Result.Batch, error) {
	var taskDb = Model.TaskDb{}
	var batchesDb = []Model.BatchDb{}
	var batches = []Result.Batch{}
	err := adapter.Orm.First(&taskDb, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	err = adapter.Orm.
		Table("batches").
		Where("task_id = ?", taskDb.ID).
		Find(&batchesDb).
		Error
	if err != nil {
		return nil, Error2.NewTaskNotFoundError()
	}

	for _, batchDb := range batchesDb {
		batch, err := batchDb.ToDomainV2()
		if err != nil {
			return nil, Error.NewRepositoryError(err.Error())
		}
		batches = append(batches, batch)
	}
	return batches, nil
}
