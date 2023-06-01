package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type FindBatchResultsAdapter struct {
	Orm *gorm.DB
}

func (adapter FindBatchResultsAdapter) Find(id Result.BatchId) ([]Result.Result, error) {
	var batchDb = Model.BatchDb{}
	var resultsDb []Model.ResultDb
	var results []Result.Result
	err := adapter.Orm.First(&batchDb, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	err = adapter.Orm.
		Table("results").
		Where("batch_id = ?", batchDb.ID).
		Find(&resultsDb).
		Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}

	for _, resultDb := range resultsDb {
		result, err := resultDb.ToDomainV2()
		if err != nil {
			return nil, Error.NewRepositoryError(err.Error())
		}
		results = append(results, result)
	}
	return results, nil
}
