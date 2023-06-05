package HostTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
)

const defaultHost = "0.0.0.0"

type voMock struct {
	value string
}

func (v *voMock) GetValue() string {
	return v.value
}

func BuildDefaultMock() Host.Vo {
	return &voMock{defaultHost}
}
func BuildHostMock(value string) Host.Vo {
	return &voMock{value}
}
