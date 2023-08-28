package main

import (
	"context"
	"fmt"
	"log"
    pb "GRPC/helloworld"
	"google.golang.org/grpc"
)

func main(){
	conn,err :=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("Failed to connect: %v",err)
	}
	defer conn.Close()
    
	client := pb.NewGreeterClient(conn)

	name :="kiran"
	var age int32
	age=21
	response,err :=client.SayHello(context.Background(),&pb.HelloRequest{Name: name,Age:age})
	if err!=nil{
		log.Fatalf("Failed to call SayHello: %v",err)
	}
	fmt.Printf("Response: %s\n", response.Message)
}