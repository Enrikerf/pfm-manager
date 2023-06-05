package SearcherTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Searcher"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/RepositoryTest"
	"reflect"
	"testing"
)

func TestSearcher_Search(t *testing.T) {
	type fields struct {
		FindByRepository Repository.FindBy
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
				FindByRepository: RepositoryTest.BuildFindByMock([]Task.Task{TaskTest.BuildDefaultTaskMock()}, nil),
			},
			want:    "no error",
			wantErr: false,
		},
		{
			name: "KO",
			fields: fields{
				FindByRepository: RepositoryTest.BuildFindByMock(nil, Error.NewRepositoryError("NewRepositoryError")),
			},
			want:    "NewRepositoryError",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searcher := &Searcher.Searcher{
				FindByRepository: tt.fields.FindByRepository,
			}
			_, err := searcher.Search(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
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
