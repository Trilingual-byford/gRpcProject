package main
//
import (
	"context"
	"fmt"
	pb "gRpcProject/protoc"
	"google.golang.org/grpc"
)

func main() {
	//1.Connect to server
	dial, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err!=nil{
		fmt.Println("Connection Error")
	}
	defer dial.Close()
	fmt.Println("Configure connection")
	client := pb.NewUserInfoServiceClient(dial)
	request := &pb.UserRequest{Name: "Bob"}
	//p := new(pb.UserRequest)
	info, err := client.GetUserInfo(context.Background(), request)
	fmt.Println("Response:",info)

}