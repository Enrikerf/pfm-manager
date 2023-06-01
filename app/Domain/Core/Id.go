package Core

import "github.com/google/uuid"

type Id interface {
	GetUuid() uuid.UUID
	GetUuidString() string
}
