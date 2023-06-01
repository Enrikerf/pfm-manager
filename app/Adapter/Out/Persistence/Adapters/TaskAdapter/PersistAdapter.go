package TaskAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"gorm.io/gorm"
)

type PersistAdapter struct {
	Orm *gorm.DB
}

func (adapter PersistAdapter) Persist(task Task.Task) {
	var currentTaskMysql Model.TaskDb
	var taskValuesToUpdate = Model.TaskDb{}
	taskValuesToUpdate.FromDomainV2(task)
	err := adapter.Orm.First(&currentTaskMysql, "uuid = ?", taskValuesToUpdate.Uuid).Error
	if err != nil {
		var taskMysql = Model.TaskDb{}
		taskMysql.FromDomainV2(task)
		_ = adapter.Orm.Create(&taskMysql).Error
	}
	adapter.Orm.Model(&currentTaskMysql).Updates(taskValuesToUpdate)
}
