package CommunicationModeTest

import "github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"

func BuildDefaultMock() CommunicationMode.Mode {
	return CommunicationMode.New(CommunicationMode.Unary)
}
