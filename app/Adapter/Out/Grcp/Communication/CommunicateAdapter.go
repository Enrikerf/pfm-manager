package Communication

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/call"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task/CommunicationMode"
	"google.golang.org/grpc"
	"io"
)

type CommunicateAdapter struct {
}

func (adapter CommunicateAdapter) Communicate(task Task.Task) Result.Batch {
	options := grpc.WithInsecure()
	connection, err := grpc.Dial(task.GetHost().GetValue()+":"+task.GetPort().GetValue(), options)
	var results []Content.Content
	if err != nil {
		fmt.Printf(err.Error())
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		defer func(connection *grpc.ClientConn) {
			err := connection.Close()
			if err != nil {
			}
		}(connection)
		client := call.NewCallServiceClient(connection)
		switch task.GetCommunicationMode() {
		case CommunicationMode.Unary:
			results = adapter.doUnaryCall(task, client)
		case CommunicationMode.ServerStream:
			results = adapter.doServerStream(task, client)
		case CommunicationMode.ClientStream:
			results = adapter.doClientStream(task, client)
		case CommunicationMode.Bidirectional:
			results = adapter.doBidirectional(task, client)
		}
	}
	batch := Result.NewBatch(task.GetId())
	batch.SetResultsFromContent(results)
	return batch
}

func (adapter CommunicateAdapter) doServerStream(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	request := &call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	responseStream, err := client.CallServerStream(context.Background(), request)

	if err != nil {
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		for {
			msg, err := responseStream.Recv()
			if err == io.EOF {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
				break
			}
			if err != nil {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
				break
			}
			result, _ := Content.NewContent(msg.GetResult())
			results = append(results, result)
		}
	}
	return results
}

func (adapter CommunicateAdapter) doBidirectional(task Task.Task, client call.CallServiceClient) []Content.Content {
	var contents []Content.Content
	stream, err := client.CallBidirectional(context.Background())
	if err != nil {
		content, _ := Content.NewContent(err.Error())
		contents = append(contents, content)
	} else {
		receiveChannel := make(chan string)
		sendChannel := make(chan string)
		go sendInStream(task, stream, sendChannel)
		go receiveServerStream(stream, receiveChannel)
		for i := range receiveChannel {
			fmt.Println(i)
			content, _ := Content.NewContent(i)
			contents = append(contents, content)
		}
		for i := range sendChannel {
			fmt.Println(i)
			content, _ := Content.NewContent(i)
			contents = append(contents, content)
		}
	}

	return contents
}

func receiveServerStream(stream call.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for {
		fmt.Println("receiving")
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end receiving")
			resultsChannel <- err.Error()
			break
		}
		if err != nil {
			fmt.Println("error receiving")
			resultsChannel <- err.Error()
			break
		}
		fmt.Println("receive result")
		resultsChannel <- response.Result
	}
	fmt.Println("finish receiving")
	close(resultsChannel)
}

func sendInStream(task Task.Task, stream call.CallService_CallBidirectionalClient, resultsChannel chan string) {
	for _, step := range task.GetSteps() {
		fmt.Println("sending")
		err := stream.Send(&call.CallRequest{Step: step.GetSentence()})
		if err != nil && err != io.EOF {
			fmt.Println("error sending")
			resultsChannel <- err.Error()
			continue
		}
		resultsChannel <- step.GetSentence()
	}
	err := stream.CloseSend()
	if err != nil {
		fmt.Println("error closing sending")
		resultsChannel <- err.Error()
	}
	fmt.Println("finish sending")
	close(resultsChannel)
}

func (adapter CommunicateAdapter) doClientStream(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	stream, err := client.CallClientStream(context.Background())
	if err != nil {
		result, _ := Content.NewContent(err.Error())
		results = append(results, result)
	} else {
		for _, step := range task.GetSteps() {
			err := stream.Send(&call.CallRequest{
				Step: step.GetSentence(),
			})
			if err != nil {
				result, _ := Content.NewContent(err.Error())
				results = append(results, result)
			}
		}
		response, err := stream.CloseAndRecv()
		if err != nil {
			result, _ := Content.NewContent(err.Error())
			results = append(results, result)
		}
		result, _ := Content.NewContent(response.GetResult())
		results = append(results, result)
	}
	return results
}

func (adapter CommunicateAdapter) doUnaryCall(task Task.Task, client call.CallServiceClient) []Content.Content {
	var results []Content.Content
	var result Content.Content
	callRequest := call.CallRequest{
		Step: task.GetSteps()[0].GetSentence(),
	}
	callResponse, err := client.CallUnary(context.Background(), &callRequest)
	if err != nil {
		result, _ = Content.NewContent(err.Error())
	} else {
		result, _ = Content.NewContent(callResponse.GetResult())
	}
	return append(results, result)
}
