package CommunicationModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name    string
		args    args
		want    ExecutionMode.Mode
		wantErr bool
	}{
		{
			name: "invalid Communication mode throw unknown error",
			args: args{
				mode: "INVALID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid Manual string",
			args: args{
				mode: "MANUAL",
			},
			want:    ExecutionMode.New(ExecutionMode.Manual),
			wantErr: false,
		},
		{
			name: "Valid ServerStream string",
			args: args{
				mode: "AUTOMATIC",
			},
			want:    ExecutionMode.New(ExecutionMode.Automatic),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecutionMode.FromString(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestNew(t *testing.T) {
//	type args struct {
//		enum ExecutionMode.Enum
//	}
//	tests := []struct {
//		name string
//		args args
//		want ExecutionMode.Enum
//	}{
//		{
//			name: "Test New",
//			args: args{
//				enum: ExecutionMode.Automatic,
//			},
//			want: ExecutionMode.Automatic,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ExecutionMode.New(tt.args.enum); !reflect.DeepEqual(got.Value(), tt.want) {
//				t.Errorf("New() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
