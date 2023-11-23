package server

import (
	"awesomeProject/hello_grpc/grpchellopbs"
	"context"
	"fmt"
)

type HelloServer struct {
	grpchellopbs.UnimplementedHelloFunctionServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *grpchellopbs.NameRequest) (*grpchellopbs.AnswerRequest, error) {
	fmt.Println(ctx)
	fmt.Println(s)
	fmt.Println(&s)
	name := req.GetName()
	answer := fmt.Sprintf("Hello, %s!", name)

	fmt.Printf("server startsss")
	return &grpchellopbs.AnswerRequest{Answer: answer}, nil
}

//func RunGRPCServer() {
//	lis, err := net.Listen("tcp", ":3333")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	helloObject := &HelloServer{}
//	grpchellopbs.RegisterHelloFunctionServer(grpcServer, helloObject)
//
//	fmt.Println("Server is listening on port 3333...")
//	if err := grpcServer.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

//func Main() {
//	RunGRPCServer()
//}
