package CommunicationModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"testing"
)

func TestNewUnknownError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "communication mode not valid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommunicationMode.NewUnknownError(); !(got.Error() == tt.want) {
				t.Errorf("NewUnknownError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}
