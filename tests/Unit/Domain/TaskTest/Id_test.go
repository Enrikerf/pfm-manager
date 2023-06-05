package TaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestLoadId(t *testing.T) {
	uuidString := "3b5baf50-e6d2-450d-854b-52c16180b27f"
	t.Run("load Task Id", func(t *testing.T) {
		uuidString, _ := uuid.Parse(uuidString)
		if got := Task.LoadId(uuidString); !reflect.DeepEqual(got.GetUuidString(), uuidString.String()) {
			t.Errorf("LoadId() = %v, want %v", got.GetUuidString(), uuidString)
		}
	})
}

func TestLoadIdFromString(t *testing.T) {
	type args struct {
		uuidString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				uuidString: defaultUuid,
			},
			want:    defaultUuid,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				uuidString: invalidUuid,
			},
			want:    "not parseable ID",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task.LoadIdFromString(tt.args.uuidString)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadIdFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("NewVo() error = %v, wantErr %v", err.Error(), tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got.GetUuidString(), tt.want) {
				t.Errorf("LoadIdFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewId(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Task Id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Task.NewId()
			if reflect.TypeOf(got.GetUuid()) != reflect.TypeOf(uuid.New()) {
				t.Errorf("NewId() = %v, want %v", reflect.TypeOf(got.GetUuid()), reflect.TypeOf(uuid.New().ID()))
			}
			if reflect.TypeOf(got.GetUuidString()) != reflect.TypeOf(uuid.New().String()) {
				t.Errorf("NewId() = %v, want %v", reflect.TypeOf(got.GetUuidString()), reflect.TypeOf(uuid.New().String()))
			}
		})
	}
}
