package FinderTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Finder"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/RepositoryTest"
	"reflect"
	"testing"
)

func TestFinder_Find(t *testing.T) {
	type fields struct {
		FindRepository Repository.Find
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				FindRepository: RepositoryTest.BuildFindMock(TaskTest.BuildDefaultTaskMock(), nil),
			},
			want:    "no error",
			wantErr: false,
		},
		{
			name: "KO",
			fields: fields{
				FindRepository: RepositoryTest.BuildFindMock(nil, nil),
			},
			want:    "NewTaskNotFoundError",
			wantErr: true,
		},
		{
			name: "KO",
			fields: fields{
				FindRepository: RepositoryTest.BuildFindMock(nil, Error.NewRepositoryError("NewRepositoryError")),
			},
			want:    "NewRepositoryError",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finder := &Finder.Finder{
				FindRepository: tt.fields.FindRepository,
			}
			_, err := finder.Find(TaskTest.DefaultId())
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("NewVo() error = %v, wantErr %v", err.Error(), tt.want)
				}
				return
			}
		})
	}
}
