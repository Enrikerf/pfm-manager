package StepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestLoadId(t *testing.T) {
	uuidString := "3b5baf50-e6d2-450d-854b-52c16180b27f"
	t.Run("load Step Id", func(t *testing.T) {
		uuidString, _ := uuid.Parse(uuidString)
		if got := Step.LoadId(uuidString); !reflect.DeepEqual(got.GetUuidString(), uuidString.String()) {
			t.Errorf("LoadId() = %v, want %v", got.GetUuidString(), uuidString)
		}
	})
}

func TestNewId(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Step Id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Step.NewId()
			if reflect.TypeOf(got.GetUuid()) != reflect.TypeOf(uuid.New()) {
				t.Errorf("NewId() = %v, want %v", reflect.TypeOf(got.GetUuid()), reflect.TypeOf(uuid.New().ID()))
			}
			if reflect.TypeOf(got.GetUuidString()) != reflect.TypeOf(uuid.New().String()) {
				t.Errorf("NewId() = %v, want %v", reflect.TypeOf(got.GetUuidString()), reflect.TypeOf(uuid.New().String()))
			}
		})
	}
}
