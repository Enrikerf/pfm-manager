package ResultAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type FindBatchResultsAfterResultAdapter struct {
	Orm *gorm.DB
}

func (adapter FindBatchResultsAfterResultAdapter) Find(id Result.BatchId, resultId Result.Id) ([]Result.Result, error) {
	var batchDb = Model.BatchDb{}
	var resultDb = Model.ResultDb{}
	var resultsDb []Model.ResultDb
	var results []Result.Result
	var err error
	//TODO: cleanup
	if resultId != nil {
		err = adapter.Orm.First(&resultDb, "uuid = ?", resultId.GetUuidString()).Error
		if err != nil {
			return nil, Error.NewRepositoryError(err.Error())
		}
	}
	err = adapter.Orm.First(&batchDb, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	if resultId != nil {
		err = adapter.Orm.
			Table("results").
			Where("batch_id = ? AND id > ?", batchDb.ID, resultDb.ID).
			Limit(50).
			Find(&resultsDb).
			Error
	} else {
		err = adapter.Orm.
			Table("results").
			Where("batch_id = ?", batchDb.ID).
			Limit(50).
			Find(&resultsDb).
			Error
	}

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
