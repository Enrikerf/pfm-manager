package Communication

import (
	"context"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"google.golang.org/grpc"
	"io"
)

type ServerStreamAdapter interface {
	Setup(host Host.Vo, port Port.Vo) error
	Communicate(task Task.Task) error
	GetIterator() (Content.Content, error)
}

func NewServerStreamAdapter() ServerStreamAdapter {
	return &serverStreamAdapter{}
}

type serverStreamAdapter struct {
	connection     *grpc.ClientConn
	client         call.CallServiceClient
	responseStream call.CallService_CallServerStreamClient
}

func (serverStreamAdapter *serverStreamAdapter) Setup(host Host.Vo, port Port.Vo) error {
	options := grpc.WithInsecure()
	var err error
	serverStreamAdapter.connection, err = grpc.Dial(host.GetValue()+":"+port.GetValue(), options)
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	serverStreamAdapter.client = call.NewCallServiceClient(serverStreamAdapter.connection)
	return nil
}

func (serverStreamAdapter *serverStreamAdapter) Communicate(task Task.Task) error {
	request := &call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	var err error
	serverStreamAdapter.responseStream, err = serverStreamAdapter.client.CallServerStream(context.Background(), request)
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	return nil
}

func (serverStreamAdapter *serverStreamAdapter) GetIterator() (Content.Content, error) {
	msg, err := serverStreamAdapter.responseStream.Recv()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	content, err := Content.NewContent(msg.GetResult())
	if err != nil {
		content, err = Content.NewContent(err.Error())
	}
	return content, nil
}
