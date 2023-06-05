package HostTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
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
			name: "Test invalid host",
			args: args{
				value: "asd....",
			},
			want:    "invalid host format",
			wantErr: true,
		},
		{
			name: "Test valid host",
			args: args{
				value: "0.0.0.0",
			},
			want:    "0.0.0.0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Host.NewVo(tt.args.value)
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
