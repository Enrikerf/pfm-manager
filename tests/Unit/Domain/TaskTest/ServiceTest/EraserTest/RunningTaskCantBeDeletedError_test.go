package EraserTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Eraser"
	"reflect"
	"testing"
)

func TestNewRunningTaskCantBeDeletedError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewInvalidHostError message",
			want: "NewRunningTaskCantBeDeletedError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eraser.NewRunningTaskCantBeDeletedError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewInvalidHostError() = %v, want %v", got, tt.want)
			}
		})
	}
}
