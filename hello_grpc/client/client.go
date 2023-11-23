package client

import (
	"awesomeProject/hello_grpc/grpchellopbs"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func RunGRPCClient(name string) {
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := grpchellopbs.NewHelloFunctionClient(conn)

	response, err := client.SayHello(context.Background(), &grpchellopbs.NameRequest{Name: name})
	if err != nil {
		fmt.Printf("Error calling SayHello: %v", err)
	}

	answer := response.Answer
	fmt.Printf("Response from server: %s\n", answer)
}
