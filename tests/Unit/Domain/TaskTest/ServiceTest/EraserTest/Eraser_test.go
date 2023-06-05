package EraserTest

import (
	Error2 "github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Eraser"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/RepositoryTest"
	"testing"
)

func TestEraser_Erase(t *testing.T) {

	type fields struct {
		FindRepository   Repository.Find
		DeleteRepository Repository.Delete
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "1-OK",
			fields: fields{
				FindRepository:   RepositoryTest.BuildFindMock(TaskTest.BuildDefaultTaskMock(), nil),
				DeleteRepository: RepositoryTest.BuildDeleteMock(nil),
			},
			wantErr: false,
		},
		{
			name: "2-Error running",
			fields: fields{
				FindRepository:   RepositoryTest.BuildFindMock(TaskTest.Get().WithStatus(Status.New(Status.Running)).Build(), nil),
				DeleteRepository: RepositoryTest.BuildDeleteMock(nil),
			},
			wantErr: true,
		},
		{
			name: "3-Error",
			fields: fields{
				FindRepository:   RepositoryTest.BuildFindMock(nil, Error2.NewRepositoryError("")),
				DeleteRepository: RepositoryTest.BuildDeleteMock(nil),
			},
			wantErr: true,
		},
		{
			name: "4-Error delete",
			fields: fields{
				FindRepository:   RepositoryTest.BuildFindMock(TaskTest.BuildDefaultTaskMock(), nil),
				DeleteRepository: RepositoryTest.BuildDeleteMock(Error2.NewRepositoryError("")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eraser := &Eraser.Eraser{
				FindRepository:   tt.fields.FindRepository,
				DeleteRepository: tt.fields.DeleteRepository,
			}
			if err := eraser.Erase(TaskTest.DefaultId()); (err != nil) != tt.wantErr {
				t.Errorf("Erase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
