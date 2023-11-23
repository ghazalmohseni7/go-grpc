package utils

import "awesomeProject/hello_grpc/client"

func Add(a, b int) int {
	client.RunGRPCClient("SUMMMMMMMMMMMMMMMM")
	return a + b
}
