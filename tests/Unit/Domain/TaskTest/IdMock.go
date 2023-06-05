package TaskTest

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task"

const defaultUuid = "4614d29e-f38e-4fb6-82e1-9b4a023195bf"
const invalidUuid = "invalid"

func DefaultId() Task.Id {
	id, _ := Task.LoadIdFromString(defaultUuid)
	return id
}
