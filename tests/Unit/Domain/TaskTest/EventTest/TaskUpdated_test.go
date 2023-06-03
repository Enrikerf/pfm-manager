package EventTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Event"
	"reflect"
	"testing"
)

func TestNewTaskUpdated(t *testing.T) {
	type args struct {
		entityId string
	}
	tests := []struct {
		name     string
		args     args
		wantId   string
		wantName string
	}{
		{
			name: "Test NewTaskUpdated Event",
			args: args{
				entityId: "id",
			},
			wantId:   "id",
			wantName: Event.TaskUpdatedEventName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Event.NewTaskUpdated(tt.args.entityId); !reflect.DeepEqual(got.GetEntityId(), tt.wantId) {
				t.Errorf("NewTaskUpdated().id = %v, want %v", got.GetEntityId(), tt.wantId)
			}
			if got := Event.NewTaskUpdated(tt.args.entityId); !reflect.DeepEqual(got.GetName(), tt.wantName) {
				t.Errorf("NewTaskUpdated().name = %v, want %v", got.GetName(), tt.wantName)
			}
		})
	}
}
