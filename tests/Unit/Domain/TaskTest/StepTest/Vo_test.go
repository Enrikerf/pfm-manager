package StepTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"reflect"
	"testing"
)

func TestNewVo(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name    string
		args    args
		want    Step.Vo
		wantErr bool
	}{
		{
			name: "",
			args: args{
				sentence: "tooLongSentence-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Step.NewVo(tt.args.sentence)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
