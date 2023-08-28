package main

import (
	pb "GRPC/task"
	"context"
	"fmt"
	"net"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

// type server struct{
// 	hw.UnimplementedGreeterServer
// }

// func(s*server) SayHello(ctx context.Context,req*hw.HelloRequest)(*hw.HelloResponse,error){
// 	return &hw.HelloResponse{
// 		Message:fmt.Sprintf("Hello,%s %v!",req.Name,req.Age),
// 	},nil
// }

type taskServiceServer struct{
	mu sync.Mutex
	tasks map[string]*pb.Task
	pb.UnimplementedTaskServiceServer
}

func(s * taskServiceServer)AddTask(ctx context.Context,req*pb.Task)(*pb.TaskResponse,error){
	s.mu.Lock()
	defer s.mu.Unlock()
	taskID:=generateID()
	req.Id=taskID
	s.tasks[taskID]=req

	return &pb.TaskResponse{Id:taskID},nil
}

func(s* taskServiceServer)GetTasks(ctx context.Context,req*pb.Empty)(*pb.TaskList,error){
	s.mu.Lock()
	defer s.mu.Unlock()
	tasks:=make([]*pb.Task,0,len(s.tasks))
	for _,task:=range s.tasks{
		tasks=append(tasks,task)
	}
	return &pb.TaskList{Tasks:tasks},nil
}

func generateID() string{
	id :=primitive.NewObjectID()
	taskID:=id.Hex()
	return taskID
}
func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		fmt.Printf("Failed to listen:%v",err)
		return
	}
	server:=grpc.NewServer()
	pb.RegisterTaskServiceServer(server,&taskServiceServer{
		tasks:make(map[string]*pb.Task),
	})
	fmt.Println("Server listening on :50051")
	if err:=server.Serve(lis); err!=nil{
		fmt.Printf("Failed to server: %v",err)
	}
}