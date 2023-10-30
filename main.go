package main

import (
	"fmt"
	"net"

	pb "github.com/mhmmdFsl/go-grpc-example/proto"
	"google.golang.org/grpc"
)

type CatServer struct {
	pb.UnimplementedTodoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Printf("falied listen to port :3000")
	}

	fmt.Printf("listening at %v", lis.Addr())

	server := grpc.NewServer()

	pb.RegisterTodoServiceServer(server, CatServer{})
	if err := server.Serve(lis); err != nil {
		fmt.Printf("grpc server failed: %v", err.Error())
	}

}
