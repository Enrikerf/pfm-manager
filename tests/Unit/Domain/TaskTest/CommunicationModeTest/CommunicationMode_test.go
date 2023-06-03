package CommunicationModeTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
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
		want    CommunicationMode.Mode
		wantErr bool
	}{
		{
			name: "invalid Communication mode throw unknow error",
			args: args{
				mode: "INVALID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid Unary string",
			args: args{
				mode: "UNARY",
			},
			want:    CommunicationMode.New(CommunicationMode.Unary),
			wantErr: false,
		},
		{
			name: "Valid ServerStream string",
			args: args{
				mode: "SERVER_STREAM",
			},
			want:    CommunicationMode.New(CommunicationMode.ServerStream),
			wantErr: false,
		},
		{
			name: "Valid ClientStream string",
			args: args{
				mode: "CLIENT_STREAM",
			},
			want:    CommunicationMode.New(CommunicationMode.ClientStream),
			wantErr: false,
		},
		{
			name: "Valid Bidirectional string",
			args: args{
				mode: "BIDIRECTIONAL",
			},
			want:    CommunicationMode.New(CommunicationMode.Bidirectional),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommunicationMode.FromString(tt.args.mode)
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
		enum CommunicationMode.Enum
	}
	tests := []struct {
		name string
		args args
		want CommunicationMode.Mode
	}{
		{
			name: "Test New",
			args: args{
				enum: CommunicationMode.Unary,
			},
			want: CommunicationMode.New(CommunicationMode.Unary),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommunicationMode.New(tt.args.enum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
