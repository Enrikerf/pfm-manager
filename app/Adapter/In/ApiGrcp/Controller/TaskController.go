package Controller

import (
	"context"
	"fmt"
	taskProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/ExecutionMode"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskController struct {
	SaveTaskUseCase   CreateTask.UseCase
	ListTasksUseCase  ListTasks.UseCase
	ReadTaskUseCase   ReadTask.UseCase
	DeleteTaskUseCase DeleteTask.UseCase
	UpdateTaskUseCase UpdateTask.UseCase
	taskProto.UnimplementedTaskServiceServer
}

func (controller TaskController) CreateTask(
	ctx context.Context,
	request *taskProto.CreateTaskRequest,
) (*taskProto.CreateTaskResponse, error) {
	protoTask := request.GetTaskParams()
	var command CreateTask.Command
	command.Host = protoTask.GetHost()
	command.Port = protoTask.GetPort()
	command.CommandSentences = protoTask.GetCommands()
	command.CommunicationMode = protoTask.GetMode().String()
	command.ExecutionMode = protoTask.GetExecutionMode().String()

	task, err := controller.SaveTaskUseCase.Create(command)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	var commandNames []string
	for _, command := range task.GetSteps() {
		commandNames = append(commandNames, command.GetSentence())
	}
	newTask := taskProto.Task{
		Uuid:          task.GetId().GetUuidString(),
		Host:          task.GetHost().GetValue(),
		Port:          task.GetPort().GetValue(),
		Commands:      commandNames,
		Mode:          string(task.GetCommunicationMode().Value()),
		Status:        string(task.GetStatus().Value()),
		ExecutionMode: string(task.GetExecutionMode().Value()),
	}
	return &taskProto.CreateTaskResponse{Task: &newTask}, nil
}

func (controller TaskController) ReadTask(
	ctx context.Context,
	request *taskProto.ReadTaskRequest,
) (*taskProto.ReadTaskResponse, error) {
	var query = ReadTask.Query{Uuid: request.GetTaskUuid()}
	task, err := controller.ReadTaskUseCase.Read(query)
	if err != nil {
		return &taskProto.ReadTaskResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("error"),
		)
	}

	return &taskProto.ReadTaskResponse{Task: &taskProto.Task{
		Uuid:          task.GetId().GetUuidString(),
		Host:          task.GetHost().GetValue(),
		Port:          task.GetPort().GetValue(),
		Commands:      nil,
		Mode:          string(task.GetCommunicationMode().Value()),
		Status:        string(task.GetStatus().Value()),
		ExecutionMode: string(task.GetExecutionMode().Value()),
	}}, nil
}

func (controller TaskController) UpdateTask(ctx context.Context, request *taskProto.UpdateTaskRequest) (*taskProto.UpdateTaskResponse, error) {

	params := request.GetParams()
	var host Host.Vo
	var port Port.Vo
	var communicationMode CommunicationMode.Mode
	var executionMode ExecutionMode.Mode
	var s Status.Status
	if params.GetHost() == nil {
		host, _ = Host.NewVo(params.GetHost().GetValue())
	}
	if params.GetPort() == nil {
		port, _ = Port.NewVo(params.GetPort().GetValue())
	}
	if params.GetStatus() != 0 {
		s, _ = Status.FromString(params.GetStatus().String())
	}
	err := controller.UpdateTaskUseCase.Update(UpdateTask.Command{
		Uuid:              request.GetTaskUuid(),
		Host:              host,
		Port:              port,
		CommunicationMode: communicationMode,
		ExecutionMode:     executionMode,
		Status:            s,
	})
	return &taskProto.UpdateTaskResponse{}, err
}

func (controller TaskController) DeleteTask(ctx context.Context, request *taskProto.DeleteTaskRequest) (*taskProto.DeleteTaskResponse, error) {
	var command = DeleteTask.Command{Uuid: request.GetTaskUuid()}
	err := controller.DeleteTaskUseCase.Delete(command)
	if err != nil {
		return &taskProto.DeleteTaskResponse{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}

	return &taskProto.DeleteTaskResponse{}, nil
}

func (controller TaskController) ListTasks(ctx context.Context, in *taskProto.ListTasksRequest) (*taskProto.ListTasksResponse, error) {
	tasks, err := controller.ListTasksUseCase.List(ListTasks.Query{})
	if err != nil {
		return &taskProto.ListTasksResponse{}, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}
	var tasksProtoArray []*taskProto.Task
	for _, task := range tasks {
		var commands []string
		if len(task.GetSteps()) > 0 {
			for _, step := range task.GetSteps() {
				commands = append(commands, step.GetSentence())
			}
		}
		t := taskProto.Task{
			Uuid:          task.GetId().GetUuidString(),
			Host:          task.GetHost().GetValue(),
			Port:          task.GetPort().GetValue(),
			Commands:      commands,
			Mode:          string(task.GetCommunicationMode().Value()),
			Status:        string(task.GetStatus().Value()),
			ExecutionMode: string(task.GetExecutionMode().Value()),
		}
		tasksProtoArray = append(tasksProtoArray, &t)
	}
	return &taskProto.ListTasksResponse{Tasks: tasksProtoArray}, nil
}
