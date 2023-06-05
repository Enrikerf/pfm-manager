package CreatorTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Event"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Repository"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Creator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	EventTest "github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/EventTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/CommunicationModeTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/ExecutionModeTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/HostTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/PortTest"
	RepositoryTest "github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/RepositoryTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/StepTest"
	"reflect"
	"testing"
)

func TestCreator_Create(t *testing.T) {
	type fields struct {
		SaveRepository Repository.Save
		Dispatcher     Event.Dispatcher
	}
	type args struct {
		host              Host.Vo
		port              Port.Vo
		stepVos           []Step.Vo
		communicationMode CommunicationMode.Mode
		executionMode     ExecutionMode.Mode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Error",
			fields: fields{
				SaveRepository: RepositoryTest.BuildSaveMock(),
				Dispatcher:     EventTest.BuildDispatcherMock(),
			},
			args: args{
				host:              HostTest.BuildDefaultMock(),
				port:              PortTest.BuildDefaultMock(),
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationModeTest.BuildDefaultMock(),
				executionMode:     ExecutionModeTest.BuildDefaultMock(),
			},
			want:    "NewTaskMustHaveAtLeastOneStepError",
			wantErr: true,
		},
		{
			name: "OK",
			fields: fields{
				SaveRepository: RepositoryTest.BuildSaveMock(),
				Dispatcher:     EventTest.BuildDispatcherMock(),
			},
			args: args{
				host:              HostTest.BuildDefaultMock(),
				port:              PortTest.BuildDefaultMock(),
				stepVos:           []Step.Vo{StepTest.BuildDefaultStepVo()},
				communicationMode: CommunicationModeTest.BuildDefaultMock(),
				executionMode:     ExecutionModeTest.BuildDefaultMock(),
			},
			want:    "no error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creator := &Creator.Creator{
				SaveRepository: tt.fields.SaveRepository,
				Dispatcher:     tt.fields.Dispatcher,
			}
			got, err := creator.Create(tt.args.host, tt.args.port, tt.args.stepVos, tt.args.communicationMode, tt.args.executionMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("NewVo() error = %v, wantErr %v", err.Error(), tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got.GetStatus().Value(), Status.Pending) {
				t.Errorf("Create() got = %v, want %v", got.GetStatus().Value(), Status.Pending)
			}
		})
	}
}
