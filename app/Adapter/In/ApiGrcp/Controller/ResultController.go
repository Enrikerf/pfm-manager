package Controller

import (
	"context"
	"fmt"
	resultProto "github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Communication/CommunicateTaskManually"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetBatchResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetTaskBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"time"
)

type ResultController struct {
	CreateBatchAndFillUseCase CommunicateTaskManually.UseCase
	GetTaskBatchesUseCase     GetTaskBatches.UseCase
	GetBatchResultsUseCase    GetBatchResults.UseCase
	StreamResultsUseCase      StreamResults.UseCase
	resultProto.UnimplementedResultServiceServer
}

func (controller ResultController) CommunicateTaskManually(
	ctx context.Context,
	request *resultProto.CommunicateTaskManuallyRequest,
) (*resultProto.CommunicateTaskManuallyResponse, error) {
	var command CommunicateTaskManually.Command
	command.TaskUuid = request.GetTaskUuid()
	batch, err := controller.CreateBatchAndFillUseCase.Communicate(command)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &resultProto.CommunicateTaskManuallyResponse{
		BatchUuid: batch.GetId().GetUuidString(),
	}, nil
}

func (controller ResultController) GetBatchResults(
	ctx context.Context,
	request *resultProto.GetBatchResultsRequest,
) (*resultProto.ListResultsResponse, error) {
	batchId, err := Result.LoadBatchIdFromString(request.GetBatchUuid())
	if err != nil {
		return nil, err
	}
	results := controller.GetBatchResultsUseCase.List(GetBatchResults.Query{BatchId: batchId})
	var resultProtoArray []*resultProto.Result
	for _, result := range results {
		resultProtoArray = append(resultProtoArray, &resultProto.Result{
			Uuid:      result.GetId().GetUuidString(),
			BatchUuid: result.GetBatchId().GetUuidString(),
			Content:   result.GetContent().GetValue(),
			CreatedAt: result.GetCreateAt().Format(time.RFC3339),
		})
	}
	return &resultProto.ListResultsResponse{Results: resultProtoArray}, nil
}

func (controller ResultController) GetTaskBatches(
	ctx context.Context,
	request *resultProto.GetTaskBatchesRequest,
) (*resultProto.ListBatchesResponse, error) {
	taskId, err := Task.LoadIdFromString(request.GetTaskUuid())
	if err != nil {
		return nil, err
	}
	batches, err := controller.GetTaskBatchesUseCase.List(GetTaskBatches.Query{TaskId: taskId})
	if err != nil {
		return nil, err
	}
	var batchProtoArray []*resultProto.Batch
	for _, batch := range batches {
		batchProtoArray = append(batchProtoArray, &resultProto.Batch{
			Uuid:      batch.GetId().GetUuidString(),
			TaskUuid:  batch.GetTaskId().GetUuidString(),
			CreatedAt: batch.GetCreatedAt().Format(time.RFC3339),
		})
	}
	return &resultProto.ListBatchesResponse{Batches: batchProtoArray}, nil
}

func (controller ResultController) StreamResults(
	request *resultProto.StreamResultsRequest,
	stream resultProto.ResultService_StreamResultsServer,
) error {
	fmt.Printf("streaming results %v\n", request)
	var lastId Result.Id
	batchUuid, err := Result.LoadBatchIdFromString(request.GetBatchUuid())
	if err != nil {
		return err
	}
	for {
		results, err := controller.StreamResultsUseCase.Stream(StreamResults.Query{
			BatchUuid: batchUuid,
			LastId:    lastId,
		})
		if err != nil {
			if err.Error() == "EndOfStreamError" {
				fmt.Printf("Task Done")
				return nil
			}
			fmt.Printf(err.Error())
			return err
		}
		if len(results) > 0 {
			var resultProtoArray []*resultProto.Result
			for _, result := range results {
				resultProtoArray = append(resultProtoArray, &resultProto.Result{
					Uuid:      result.GetId().GetUuidString(),
					BatchUuid: result.GetBatchId().GetUuidString(),
					Content:   result.GetContent().GetValue(),
					CreatedAt: result.GetCreateAt().Format(time.RFC3339),
				})
			}
			err := stream.Send(&resultProto.StreamResultsResponse{Results: resultProtoArray})
			if err != nil {
				return err
			}
			lastId = results[len(results)-1].GetId()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
