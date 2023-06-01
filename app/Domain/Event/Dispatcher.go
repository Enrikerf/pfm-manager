package Event

type Dispatcher interface {
	Dispatch(event Event)
}
