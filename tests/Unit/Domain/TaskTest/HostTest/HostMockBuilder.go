package HostTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
)

type voMock struct {
	value string
}

func (v *voMock) GetValue() string {
	return v.value
}

func NewVoMock(value string) Host.Vo {
	return &voMock{value}
}
