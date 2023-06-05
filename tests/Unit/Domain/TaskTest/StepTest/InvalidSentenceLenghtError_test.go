package StepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"reflect"
	"testing"
)

func TestNewInvalidSentenceLengthError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewUnknownError message",
			want: "step.sentence length must be less than 255",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Step.NewInvalidSentenceLengthError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewInvalidSentenceLengthError() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}
