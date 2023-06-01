package Event

type Event interface {
	GetName() string
	GetEntityId() string
}
