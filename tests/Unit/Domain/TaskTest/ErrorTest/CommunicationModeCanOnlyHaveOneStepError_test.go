package ErrorTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"reflect"
	"testing"
)

func TestNewCommunicationModeCanOnlyHaveOneStepError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "NewCommunicationModeCanOnlyHaveOneStepError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error.NewCommunicationModeCanOnlyHaveOneStepError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewCommunicationModeCanOnlyHaveOneStepError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}

func TestNewManualBidirectionalTaskOnlyCanHave2StepsError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "NewManualBidirectionalTaskOnlyCanHave2StepsError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error.NewManualBidirectionalTaskOnlyCanHave2StepsError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewManualBidirectionalTaskOnlyCanHave2StepsError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}

func TestNewTaskMustHaveAtLeastOneStepError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "NewTaskMustHaveAtLeastOneStepError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error.NewTaskMustHaveAtLeastOneStepError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewTaskMustHaveAtLeastOneStepError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}

func TestNewTaskNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "NewTaskNotFoundError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error.NewTaskNotFoundError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewTaskNotFoundError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}
