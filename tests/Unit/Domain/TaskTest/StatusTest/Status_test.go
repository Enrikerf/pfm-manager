package StatusTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
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
		want    Status.Status
		wantErr bool
	}{
		{
			name: "Invalid Status throw unknown error",
			args: args{
				mode: "INVALID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid Pending string",
			args: args{
				mode: "PENDING",
			},
			want:    Status.New(Status.Pending),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Status.FromString(tt.args.mode)
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

func TestNew(t *testing.T) {
	type args struct {
		enum Status.Enum
	}
	tests := []struct {
		name string
		args args
		want Status.Status
	}{
		{
			name: "Test New",
			args: args{
				enum: Status.Pending,
			},
			want: Status.New(Status.Pending),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Status.New(tt.args.enum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
