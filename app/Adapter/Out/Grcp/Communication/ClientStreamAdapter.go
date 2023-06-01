package Communication

import (
	"context"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"google.golang.org/grpc"
)

type ClientStreamAdapter interface {
	Communicate(task Task.Task) (Content.Content, error)
}

func NewClientStreamAdapter() ClientStreamAdapter {
	return &clientStreamAdapter{}
}

type clientStreamAdapter struct {
}

func (a *clientStreamAdapter) Communicate(task Task.Task) (Content.Content, error) {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.GetHost().GetValue()+":"+task.GetPort().GetValue(), options)
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	client := call.NewCallServiceClient(connection)
	stream, err := client.CallClientStream(context.Background())
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	for _, step := range task.GetSteps() {
		err := stream.Send(&call.CallRequest{Step: step.GetSentence()})
		if err != nil {
			return nil, Error.NewRepositoryError(err.Error())
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	result, err := Content.NewContent(response.GetResult())
	if err != nil {
		return nil, Error.NewRepositoryError(err.Error())
	}
	return result, nil

}
