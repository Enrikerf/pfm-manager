package TaskAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"gorm.io/gorm"
)

type DeleteAdapter struct {
	Orm *gorm.DB
}

func (adapter DeleteAdapter) Delete(id Task.Id) error {
	var taskMysql = Model.TaskDb{}
	err := adapter.Orm.Delete(&taskMysql, "uuid = ?", id.GetUuidString()).Error
	if err != nil {
		return err
	}
	return nil
}
