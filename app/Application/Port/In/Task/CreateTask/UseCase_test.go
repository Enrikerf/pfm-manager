package CreateTask

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Service/Creator"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/EventTest"
	"github.com/Enrikerf/pfm/commandManager/tests/Unit/Domain/TaskTest/RepositoryTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_useCase_Create(t *testing.T) {
	type fields struct {
		creator Creator.Creator
	}
	type args struct {
		command Command
	}
	const hostString = "X.X.X.X"
	const portString = "8000"
	const sentenceString = "sentence"
	host, _ := Host.NewVo(hostString)
	port, _ := Port.NewVo(portString)
	step, _ := Step.NewVo(sentenceString)
	var stepVos []Step.Vo
	stepVos = append(stepVos, step)
	task, _ := Task.New(host, port, stepVos, CommunicationMode.Unary, ExecutionMode.Manual)
	creator := Creator.Creator{
		SaveRepository: RepositoryTest.SaveMock{},
		Dispatcher:     EventTest.DispatcherMock{},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Task.Task
		wantErr bool
	}{
		{
			"x",
			fields{
				creator: creator,
			},
			args{
				command: Command{
					Host:              hostString,
					Port:              portString,
					CommandSentences:  []string{sentenceString},
					CommunicationMode: "UNARY",
					ExecutionMode:     "MANUAL",
				},
			},
			task,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := &useCase{
				creator: tt.fields.creator,
			}
			_, err := useCase.Create(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Nil(t, err)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Create() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
