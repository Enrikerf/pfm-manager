package Communication

import (
	"context"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Core/Error"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Host"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Port"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/Step"
	"google.golang.org/grpc"
	"io"
)

type BidirectionalAdapter interface {
	Setup(host Host.Vo, port Port.Vo) error
	Write(step Step.Step) error
	Read() (Content.Content, error)
	Close() error
}

func NewBidirectionalAdapter() BidirectionalAdapter {
	return &bidirectionalAdapter{}
}

type bidirectionalAdapter struct {
	connection *grpc.ClientConn
	client     call.CallService_CallBidirectionalClient
}

func (manualAdapter *bidirectionalAdapter) Setup(host Host.Vo, port Port.Vo) error {
	options := grpc.WithInsecure()
	var err error
	manualAdapter.connection, err = grpc.Dial(host.GetValue()+":"+port.GetValue(), options)
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	client := call.NewCallServiceClient(manualAdapter.connection)
	manualAdapter.client, err = client.CallBidirectional(context.Background())
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	return nil
}
func (manualAdapter *bidirectionalAdapter) Write(step Step.Step) error {
	err := manualAdapter.client.Send(&call.CallRequest{
		Step: step.GetSentence(),
	})
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	return nil
}
func (manualAdapter *bidirectionalAdapter) Read() (Content.Content, error) {
	response, err := manualAdapter.client.Recv()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return Content.NewContent(response.Result)
}
func (manualAdapter *bidirectionalAdapter) Close() error {
	err := manualAdapter.connection.Close()
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	err = manualAdapter.client.CloseSend()
	if err != nil {
		return Error.NewRepositoryError(err.Error())
	}
	return nil
}
