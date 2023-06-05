package HostTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"reflect"
	"testing"
)

func TestNewInvalidHostError(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check NewInvalidHostError message",
			want: "invalid host format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Host.NewInvalidHostError(); !reflect.DeepEqual(got.Error(), tt.want) {
				t.Errorf("NewInvalidHostError() = %v, want %v", got, tt.want)
			}
		})
	}
}
