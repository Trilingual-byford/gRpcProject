package main

import (
	"context"
	"fmt"
	pb "gRpcProject/protoc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserInfoService struct {}
var u=UserInfoService{}
func (s *UserInfoService)GetUserInfo(ctx context.Context,request *pb.UserRequest)(response *pb.UserResponse,err error){
	//Search UserInfo with UserName
	name := request.Name
	log.Println(name)
	userResponse := &pb.UserResponse{Id: 1, Name: "Lucy", Age: 33}
	return userResponse, nil
}
func main() {
	s := "localhost:9090"
	listen, err := net.Listen("tcp", s)
	if err!=nil{
		fmt.Errorf("Listen Error!")
	}
	server := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(server,&u)

	//server.Serve(listen)
	server.Serve(listen)
}