package ResultAdapter

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type PersistAdapter struct {
	Orm *gorm.DB
}

func (adapter PersistAdapter) Persist(result Result.Result) {
	var currentMysqlModel Model.ResultDb
	var modelToUpdate = Model.ResultDb{}
	err := adapter.Orm.First(&currentMysqlModel, "uuid = ?", modelToUpdate.Uuid).Error
	if err != nil {
		var batchDb Model.BatchDb
		err := adapter.Orm.First(&batchDb, "uuid = ?", result.GetBatchId().GetUuidString()).Error
		if err != nil {
			fmt.Println(err.Error())
		}
		modelToUpdate.FromDomainV2(result)
		modelToUpdate.BatchID = batchDb.ID
		err = adapter.Orm.Create(&modelToUpdate).Error
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		adapter.Orm.Model(&currentMysqlModel).Updates(modelToUpdate)
	}
}
