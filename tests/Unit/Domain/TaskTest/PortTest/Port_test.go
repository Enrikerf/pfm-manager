package PortTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"reflect"
	"testing"
)

func TestNewVo(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test invalid port",
			args: args{
				value: "asd....",
			},
			want:    "invalid port format",
			wantErr: true,
		},
		{
			name: "Test valid port",
			args: args{
				value: "8080",
			},
			want:    "8080",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Port.NewVo(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("NewVo() error = %v, wantErr %v", err.Error(), tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got.GetValue(), tt.want) {
				t.Errorf("NewVo() got = %v, want %v", got.GetValue(), tt.want)
			}
		})
	}
}
