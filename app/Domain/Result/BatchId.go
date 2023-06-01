package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/google/uuid"
)

type BatchId interface {
	Core.Id
}

type batchId struct {
	uuid uuid.UUID
}

func NewBatchId() BatchId {
	id := &batchId{}
	id.uuid = uuid.New()
	return id
}

func LoadBatchId(uuid uuid.UUID) BatchId {
	id := &batchId{}
	id.uuid = uuid
	return id
}

func LoadBatchIdFromString(uuidString string) (BatchId, error) {
	parse, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, Error.NewInvalidUuidError()
	}
	return LoadBatchId(parse), nil
}

func (id batchId) GetUuid() uuid.UUID {
	return id.uuid
}

func (id batchId) GetUuidString() string {
	return id.uuid.String()
}
