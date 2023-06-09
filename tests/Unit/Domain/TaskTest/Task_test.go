package TaskTest

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	defaultHost := "0.0.0.0"
	defaultPort := "8080"
	defaultSentence := "Sentence"
	host, _ := Host.NewVo(defaultHost)
	port, _ := Port.NewVo(defaultPort)
	step, _ := Step.NewVo(defaultSentence)
	id := Task.NewId()
	type args struct {
		id                Task.Id
		host              Host.Vo
		port              Port.Vo
		stepVos           []Step.Vo
		communicationMode CommunicationMode.Mode
		executionMode     ExecutionMode.Mode
		status            Status.Status
	}
	tests := []struct {
		name    string
		args    args
		want    Status.Status
		wantErr bool
	}{
		{
			name: "load OK",
			args: args{
				id:                id,
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
				status:            Status.New(Status.Done),
			},
			want:    Status.New(Status.Done),
			wantErr: false,
		},
		{
			name: "load KO",
			args: args{
				id:                id,
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
				status:            Status.New(Status.Done),
			},
			want:    Status.New(Status.Done),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task.Load(tt.args.id, tt.args.host, tt.args.port, tt.args.stepVos, tt.args.communicationMode, tt.args.executionMode, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.GetStatus(), tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	defaultHost := "0.0.0.0"
	defaultPort := "8080"
	defaultSentence := "Sentence"
	host, _ := Host.NewVo(defaultHost)
	port, _ := Port.NewVo(defaultPort)
	step, _ := Step.NewVo(defaultSentence)
	type args struct {
		host              Host.Vo
		port              Port.Vo
		stepVos           []Step.Vo
		communicationMode CommunicationMode.Mode
		executionMode     ExecutionMode.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1-NewTaskMustHaveAtLeastOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    Error.NewTaskMustHaveAtLeastOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "2-OK",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step},
				communicationMode: CommunicationMode.New(CommunicationMode.ClientStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    "No error",
			wantErr: false,
		},
		{
			name: "3-OK",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step},
				communicationMode: CommunicationMode.New(CommunicationMode.Bidirectional),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    "No error",
			wantErr: false,
		},
		{
			name: "4-OK",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step},
				communicationMode: CommunicationMode.New(CommunicationMode.Unary),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    "No error",
			wantErr: false,
		},
		{
			name: "5-OK",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    "No error",
			wantErr: false,
		},
		{
			name: "6 CommunicationModeCanOnlyHaveOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step, step, step},
				communicationMode: CommunicationMode.New(CommunicationMode.Unary),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    Error.NewCommunicationModeCanOnlyHaveOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "7 CommunicationModeCanOnlyHaveOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step, step, step},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    Error.NewCommunicationModeCanOnlyHaveOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "8-OK",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step, step, step},
				communicationMode: CommunicationMode.New(CommunicationMode.ClientStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    "No error",
			wantErr: false,
		},
		{
			name: "9-ManualBidirectionalTaskOnlyCanHave2StepsError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{step, step, step},
				communicationMode: CommunicationMode.New(CommunicationMode.Bidirectional),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    Error.NewManualBidirectionalTaskOnlyCanHave2StepsError().Error(),
			wantErr: true,
		},
		{
			name: "10-TaskMustHaveAtLeastOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.ClientStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    Error.NewTaskMustHaveAtLeastOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "11-TaskMustHaveAtLeastOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.Bidirectional),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    Error.NewTaskMustHaveAtLeastOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "12-TaskMustHaveAtLeastOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.Unary),
				executionMode:     ExecutionMode.New(ExecutionMode.Automatic),
			},
			want:    Error.NewTaskMustHaveAtLeastOneStepError().Error(),
			wantErr: true,
		},
		{
			name: "13-TaskMustHaveAtLeastOneStepError",
			args: args{
				host:              host,
				port:              port,
				stepVos:           []Step.Vo{},
				communicationMode: CommunicationMode.New(CommunicationMode.ServerStream),
				executionMode:     ExecutionMode.New(ExecutionMode.Manual),
			},
			want:    Error.NewTaskMustHaveAtLeastOneStepError().Error(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task.New(tt.args.host, tt.args.port, tt.args.stepVos, tt.args.communicationMode, tt.args.executionMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.want) {
					t.Errorf("NewVo() error = %v, wantErr %v", err.Error(), tt.want)
				}
				return
			}
			if reflect.TypeOf(got.GetId().GetUuid()) != reflect.TypeOf(uuid.New()) {
				t.Errorf("NewVo() error = %v, wantErr %v", reflect.TypeOf(got.GetId()), reflect.TypeOf(uuid.New()))
			}
			if !reflect.DeepEqual(got.GetHost().GetValue(), defaultHost) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetHost().GetValue(), defaultHost) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetPort().GetValue(), defaultPort) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetSteps()[0].GetSentence(), defaultSentence) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetCommunicationMode(), tt.args.communicationMode) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetExecutionMode(), tt.args.executionMode) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetStatus().Value(), Status.Pending) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task_Setters(t *testing.T) {
	defaultHost := "0.0.0.0"
	defaultPort := "8080"
	defaultSentence := "Sentence"
	newHost := "0.0.0.1"
	newPort := "8081"
	host, _ := Host.NewVo(defaultHost)
	port, _ := Port.NewVo(defaultPort)
	step, _ := Step.NewVo(defaultSentence)

	t.Run("Test Task Setters", func(t *testing.T) {
		got, _ := Task.New(host, port, []Step.Vo{step}, CommunicationMode.New(CommunicationMode.Unary), ExecutionMode.New(ExecutionMode.Manual))
		newHost, _ := Host.NewVo(newHost)
		newPort, _ := Port.NewVo(newPort)
		newStatus := Status.New(Status.Done)
		got.SetHost(newHost)
		got.SetPort(newPort)
		got.SetStatus(newStatus)

		if !reflect.DeepEqual(got.GetHost(), newHost) {
			t.Errorf("New() got = %v, want %v", got.GetHost(), newHost)
		}
		if !reflect.DeepEqual(got.GetPort(), newPort) {
			t.Errorf("New() got = %v, want %v", got.GetPort(), newPort)
		}
		if !reflect.DeepEqual(got.GetStatus(), newStatus) {
			t.Errorf("New() got = %v, want %v", got, newStatus)
		}
	})

}
