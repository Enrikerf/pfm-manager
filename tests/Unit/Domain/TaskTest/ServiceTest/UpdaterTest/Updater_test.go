package UpdaterTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Updater"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"testing"
)

func TestUpdater_Update(t *testing.T) {
	type fields struct {
		FindRepository Repository.Find
		SaveRepository Repository.Save
		Dispatcher     Event.Dispatcher
	}
	type args struct {
		id     Task.Id
		host   Host.Vo
		port   Port.Vo
		status Status.Status
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updater := &Updater.Updater{
				FindRepository: tt.fields.FindRepository,
				SaveRepository: tt.fields.SaveRepository,
				Dispatcher:     tt.fields.Dispatcher,
			}
			if err := updater.Update(tt.args.id, tt.args.host, tt.args.port, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
