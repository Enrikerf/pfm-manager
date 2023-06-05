package SearcherTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Searcher"
	"reflect"
	"testing"
)

func TestSearcher_Search(t *testing.T) {
	type fields struct {
		FindByRepository Repository.FindBy
	}
	type args struct {
		conditions interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searcher := &Searcher.Searcher{
				FindByRepository: tt.fields.FindByRepository,
			}
			got, err := searcher.Search(tt.args.conditions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
