package Communication

import (
	"context"
	"errors"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"google.golang.org/grpc"
)

type UnaryAdapter interface {
	Communicate(task Task.Task) (Content.Content, error)
}

func NewUnaryAdapter() UnaryAdapter {
	return &unaryAdapter{}
}

type unaryAdapter struct {
}

func (manualAdapter *unaryAdapter) Communicate(task Task.Task) (Content.Content, error) {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.GetHost().GetValue()+":"+task.GetPort().GetValue(), options)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, errors.New("can't setup")
	}
	client := call.NewCallServiceClient(connection)
	callRequest := call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	callResponse, err := client.CallUnary(context.Background(), &callRequest)
	var result Content.Content
	if err != nil {
		result, _ = Content.NewContent(err.Error())
	} else {
		result, _ = Content.NewContent(callResponse.GetResult())
	}

	return result, nil
}
