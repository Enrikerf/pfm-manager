package HostTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
)

type voMock struct {
	value string
}

func (v *voMock) GetValue() string {
	return v.value
}

func NewVoMock(value string) Port.Vo {
	return &voMock{value}
}
