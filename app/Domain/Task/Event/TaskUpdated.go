package Event

const TaskUpdatedEventName = "task.updated"

type TaskUpdated interface {
	GetName() string
	GetEntityId() string
}

type taskUpdated struct {
	name     string
	entityId string
}

func NewTaskUpdated(entityId string) TaskUpdated {
	self := &taskUpdated{TaskUpdatedEventName, entityId}
	return self
}

func (t *taskUpdated) GetName() string {
	return t.name
}

func (t *taskUpdated) GetEntityId() string {
	return t.entityId
}
