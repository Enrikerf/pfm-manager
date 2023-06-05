package StepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	uuidString := "3b5baf50-e6d2-450d-854b-52c16180b27f"
	vo, _ := Step.NewVo("step")
	uuidFromString, _ := uuid.Parse(uuidString)
	id := Step.LoadId(uuidFromString)
	type args struct {
		id     Step.Id
		stepVo Step.Vo
	}
	tests := []struct {
		name     string
		args     args
		sentence string
	}{
		{
			name: "test Load Step",
			args: args{
				id:     id,
				stepVo: vo,
			},
			sentence: "step",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Step.Load(tt.args.id, tt.args.stepVo)
			if !reflect.DeepEqual(got.GetSentence(), tt.sentence) {
				t.Errorf("Load() = %v, want %v", got, tt.sentence)
			}
			if !reflect.DeepEqual(got.GetId().GetUuidString(), uuidString) {
				t.Errorf("Load() = %v, want %v", got.GetId().GetUuidString(), uuidString)
			}
		})
	}
}

func TestNew(t *testing.T) {
	vo, _ := Step.NewVo("step")
	type args struct {
		stepVo Step.Vo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test New Step",
			args: args{
				stepVo: vo,
			},
			want: "step",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Step.New(tt.args.stepVo); !reflect.DeepEqual(got.GetSentence(), tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
