package ResultAdapter

import (
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"gorm.io/gorm"
)

type PersistBatchAdapter struct {
	Orm *gorm.DB
}

func (adapter PersistBatchAdapter) Persist(batch Result.Batch) {
	var currentBatchDb Model.BatchDb
	var newBatchDb = Model.BatchDb{}
	err := adapter.Orm.First(&currentBatchDb, "uuid = ?", batch.GetId().GetUuidString()).Error
	if err != nil {
		var taskDb Model.TaskDb
		err := adapter.Orm.First(&taskDb, "uuid = ?", batch.GetTaskId().GetUuidString()).Error
		if err != nil {
			fmt.Println("task not found")
		}
		newBatchDb.FromDomainV2(batch)
		newBatchDb.TaskID = taskDb.ID
		err = adapter.Orm.Create(&newBatchDb).Error
		if err != nil {
			fmt.Println("batch can not be created")
		}
		for _, result := range batch.GetResults() {
			var newResultDb = Model.ResultDb{}
			newResultDb.FromDomainV2(result)
			newResultDb.BatchID = newBatchDb.ID
			err = adapter.Orm.Create(&newResultDb).Error
			if err != nil {
				fmt.Println("result can not be created")
			}
		}
	} else {
		adapter.Orm.Model(&currentBatchDb).Updates(newBatchDb)
		//TODO: Update of results not implemented
	}
}
