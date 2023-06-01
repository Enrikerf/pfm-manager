package Task

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/google/uuid"
)

type Id interface {
	Core.Id
}

type id struct {
	uuid uuid.UUID
}

func NewId() Id {
	id := &id{}
	id.uuid = uuid.New()
	return id
}

func LoadId(uuid uuid.UUID) Id {
	id := &id{}
	id.uuid = uuid
	return id
}

func LoadIdFromString(uuidString string) (Id, error) {
	parse, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, Error.NewInvalidUuidError()
	}
	return LoadId(parse), nil
}

func (id id) GetUuid() uuid.UUID {
	return id.uuid
}

func (id id) GetUuidString() string {
	return id.uuid.String()
}
