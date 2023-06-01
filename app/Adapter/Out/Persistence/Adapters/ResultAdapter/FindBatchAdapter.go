package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type FindBatchAdapter struct {
	Orm *gorm.DB
}

func (adapter FindBatchAdapter) Find(id Result.BatchId) (Result.Batch, error) {
	var batchDb = Model.BatchDb{}
	err := adapter.Orm.First(&batchDb, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	batch, err := batchDb.ToDomainV2()
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	return batch, nil

}
