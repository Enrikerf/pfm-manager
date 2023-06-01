package TaskAdapter

import (
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/Out/Persistence/Model"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"gorm.io/gorm"
)

type FindByAdapter struct {
	Orm *gorm.DB
}

func (adapter FindByAdapter) FindBy(conditions interface{}) ([]Task.Task, error) {

	var tasks []Model.TaskDb
	var domainTasks []Task.Task
	err := adapter.Orm.
		Table("tasks").
		Preload("Steps").
		Where(conditions).
		Joins("left join steps on steps.task_id = tasks.id").
		Group("tasks.id").
		Find(&tasks).
		Error
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}

	for _, task := range tasks {
		task, err := task.ToDomainV2()
		if err != nil {
			return nil, Error.NewRepositoryError(err.Error())
		}
		domainTasks = append(domainTasks, task)
	}
	return domainTasks, nil
}
