package main

import (
pb "grpc/customer"
"sync"

"golang.org/x/net/context"
)

type CustomerServiceServer struct{
mu sync.Mutex
customer map[string]*pb.CustomerDetails
pb.UnimplementedCustomerserviceServer
}

func (c * CustomerServiceServer)InsertCustomer(ctx context.Context,req * pb.CustomerDetails)(*pb.FetchDetails,error){
c.mu.Lock()
defer c.mu.Unlock()
accID:=generateID()
req.AccID=accID
c.customer[accID]=req
return &pb.FetchDetails{
Customer: []*pb.CustomerDetails{},
},nil
}
