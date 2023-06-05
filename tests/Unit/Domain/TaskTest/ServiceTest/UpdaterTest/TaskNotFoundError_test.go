package UpdaterTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Updater"
	"reflect"
	"testing"
)

func TestNewTaskNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewInvalidHostError message",
			want: "NewTaskNotFoundError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Updater.NewTaskNotFoundError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewInvalidHostError() = %v, want %v", got, tt.want)
			}
		})
	}
}
