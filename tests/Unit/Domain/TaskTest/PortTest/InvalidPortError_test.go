package PortTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"reflect"
	"testing"
)

func TestNewInvalidHostError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewInvalidPortError message",
			want: "invalid port format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Port.NewInvalidPortError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewInvalidPortError() = %v, want %v", got, tt.want)
			}
		})
	}
}
