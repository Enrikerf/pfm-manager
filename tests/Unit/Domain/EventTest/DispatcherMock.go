package EventTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
)

type DispatcherMock struct {
}

func (e DispatcherMock) Dispatch(event Event.Event) {
}
