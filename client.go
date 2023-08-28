package main

import (
	"fmt"
	"context"
	pb "GRPC/task"
	"google.golang.org/grpc"
	"log"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("Failed to connect: %v",err)
	}
	defer conn.Close()

	// client:=pb.NewGreeterClient(conn)
	client:=pb.NewTaskServiceClient(conn)

	//Add a task

	task:=&pb.Task{
		Title:"Buy groceries",
	}
	addResp,err:=client.AddTask(context.Background(),task)

	if err!=nil{
		log.Fatalf("Failed to add tasks: %v",err)
	}
	fmt.Printf("Added task with ID: %s\n",addResp.Id)




	// response,err:=client.AddTask(context.Background(),&pb.HelloRequest{Name:name})
	tasksResp,err:=client.GetTasks(context.Background(),&pb.Empty{})
	if err!=nil{
		log.Fatalf("Failed to call SayHello: %v",err)
	}
	fmt.Println("Tasks:")
	for _,task:=range tasksResp.Tasks{
          fmt.Printf("ID: %s,Title:%s,completed:%v\n",task.Id,task.Title,task.Completed)
	}
}