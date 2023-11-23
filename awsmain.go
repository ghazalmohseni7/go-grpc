package main

import (
	grpc_hello_client "awesomeProject/hello_grpc/client"
	"awesomeProject/hello_grpc/grpchellopbs"
	"awesomeProject/hello_grpc/server"
	"awesomeProject/utils"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {

	lis, errors := net.Listen("tcp", ":3333")
	if errors != nil {
		fmt.Println("failed to listen %s", errors.Error())
		return
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	helloObject := &server.HelloServer{}

	grpchellopbs.RegisterHelloFunctionServer(grpcServer, helloObject)

	go func() {
		// Start the gRPC server in a goroutine
		fmt.Println("go routinnnnnnneeee start the grpc server")
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("failed to serve %s", err.Error())
		}
	}()

	// Your other code...
	grpc_hello_client.RunGRPCClient("gombazzzzzz")
	fmt.Println(utils.Add(3, 4))
	fmt.Println(utils.Diff(20, 4))
	// Block to keep the program running
	select {}
}
