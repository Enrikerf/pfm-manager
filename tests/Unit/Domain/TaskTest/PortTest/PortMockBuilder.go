package PortTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
)

const defaultPort = "8080"

type voMock struct {
	value string
}

func (v *voMock) GetValue() string {
	return v.value
}

func BuildDefaultMock() Port.Vo {
	return &voMock{defaultPort}
}
func BuildHostMock(value string) Port.Vo {
	return &voMock{value}
}
