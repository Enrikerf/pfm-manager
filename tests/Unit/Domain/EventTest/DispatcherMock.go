package Event

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Event"

type dispatcherMock struct {
}

func (d dispatcherMock) Dispatch(event Event.Event) {

}

func BuildDispatcherMock() Event.Dispatcher {
	return dispatcherMock{}
}
