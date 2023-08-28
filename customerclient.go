package main

import (
"context"
"fmt"
pb "grpc/customer"
"log"

"google.golang.org/grpc"
)

func main(){
conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure() )
if err!=nil{ log.Fatal("failed1 %v",err)
}
defer conn.Close()
client:=pb.NewCustomerServiceClient(conn)
customer:=&pb.CustomerDetails{
Name: "",
AccID: req.AccID,
Balance: 0,
BankID: 0,
} // response,err:=client1.SayHello(context.Background(),&hw.HelloRequest{Name:name,Age:int32(age)})
addres,err:=client.AddTask(context.Background(),customer)
if err!=nil{ log.Fatal("failed2 %v",err)
}
fmt.Printf("Response of add :%s\n",addres.Id)
taskres,err:=client.GetTasks(context.Background(),&pb.Empty{})
if err!=nil{ log.Fatal("failed2 %v",err)
}
fmt.Printf("Response of get :")
for _,task:=range taskres.Tasks{
fmt.Printf("ID:%s,Title:%s,Completed:%v\n",task.Id,task.Title,task.Completed)
}}
