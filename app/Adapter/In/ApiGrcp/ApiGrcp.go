package ApiGrcp

import (
	"context"
	"fmt"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/Controller"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/result"
	"github.com/Enrikerf/pfm/commandManager/app/Adapter/In/ApiGrcp/gen/task"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Communication/CommunicateTaskManually"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetBatchResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/GetTaskBatches"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Result/StreamResults"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/CreateTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/DeleteTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ListTasks"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/ReadTask"
	"github.com/Enrikerf/pfm/commandManager/app/Application/Port/In/Task/UpdateTask"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type ApiGrpc struct {
	createTaskUseCase      CreateTask.UseCase
	listTasksUseCase       ListTasks.UseCase
	readTaskUseCase        ReadTask.UseCase
	deleteTaskUseCase      DeleteTask.UseCase
	updateTaskUseCase      UpdateTask.UseCase
	executeTaskManually    CommunicateTaskManually.UseCase
	getBatchResultsUseCase GetBatchResults.UseCase
	streamResultsUseCase   StreamResults.UseCase
	getTaskBatches         GetTaskBatches.UseCase
	serverHost             string
	serverPort             string
	grpcServer             *grpc.Server
	listener               net.Listener
}

func (api *ApiGrpc) Initialize(
	createTaskUseCase CreateTask.UseCase,
	readTaskUseCase ReadTask.UseCase,
	updateTaskUseCase UpdateTask.UseCase,
	deleteTaskUseCase DeleteTask.UseCase,
	listTasksUseCase ListTasks.UseCase,
	executeTaskManuallyUseCase CommunicateTaskManually.UseCase,
	getBatchResultsUseCase GetBatchResults.UseCase,
	streamResultsUseCase StreamResults.UseCase,
	getBatchesUseCase GetTaskBatches.UseCase,
	host string,
	port string,
) {
	fmt.Println("Starting Sentence Manager...")

	api.createTaskUseCase = createTaskUseCase
	api.listTasksUseCase = listTasksUseCase
	api.readTaskUseCase = readTaskUseCase
	api.deleteTaskUseCase = deleteTaskUseCase
	api.updateTaskUseCase = updateTaskUseCase

	api.executeTaskManually = executeTaskManuallyUseCase
	api.getBatchResultsUseCase = getBatchResultsUseCase
	api.streamResultsUseCase = streamResultsUseCase
	api.getTaskBatches = getBatchesUseCase

	api.serverHost = host
	api.serverPort = port
	api.loadServer()
	api.configControllers()
	api.loadListener()
}

func (api *ApiGrpc) Run() {
	if os.Getenv("APP_DEBUG") == "true" {
		reflection.Register(api.grpcServer)
	}
	go func() {
		fmt.Println("Starting at: " + api.serverHost + api.serverPort)
		if err := api.grpcServer.Serve(api.listener); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	// Wait for control C to exit
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	// Bock until a signal is received
	<-channel
	fmt.Println("Stopping the server")
	api.grpcServer.Stop()
	fmt.Println("closing the Listener")
	err := api.listener.Close()
	if err != nil {
		return
	}
	fmt.Println("End of program")
}

func (api *ApiGrpc) configControllers() {
	var taskController = Controller.TaskController{
		SaveTaskUseCase:   api.createTaskUseCase,
		ListTasksUseCase:  api.listTasksUseCase,
		ReadTaskUseCase:   api.readTaskUseCase,
		DeleteTaskUseCase: api.deleteTaskUseCase,
		UpdateTaskUseCase: api.updateTaskUseCase,
	}
	task.RegisterTaskServiceServer(api.grpcServer, taskController)

	var resultController = Controller.ResultController{
		CreateBatchAndFillUseCase:        api.executeTaskManually,
		GetTaskBatchesUseCase:            api.getTaskBatches,
		GetBatchResultsUseCase:           api.getBatchResultsUseCase,
		StreamResultsUseCase:             api.streamResultsUseCase,
		UnimplementedResultServiceServer: result.UnimplementedResultServiceServer{},
	}
	result.RegisterResultServiceServer(api.grpcServer, resultController)
}

func (api *ApiGrpc) loadServer() {
	//var serverOptions []grpc.ServerOption
	errHandler := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			fmt.Printf("method %q failed: %s", info.FullMethod, err)
		}
		return resp, err
	}
	api.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(errHandler))
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func (api *ApiGrpc) loadListener() {
	listener, err := net.Listen("tcp", api.serverHost+api.serverPort)
	if err != nil {
		//log.Fatalf("failed to listen at: " + api.serverHost + api.serverPort)
		log.Fatalf(err.Error())
	}
	api.listener = listener
}
