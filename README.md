# go-grpc
a simple gRPC in Golang contains commands to create protobuf files<br />

first have to install protoc: https://gist.github.com/jjeffery/6e3a4d18ffbe1fc715be403f6391c5f4<br />
make sure to set the path in windows<br />

install these libraries : <br />
go get google.golang.org/protobuf/cmd/protoc-gen-go<br />
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc<br />


create a project then create a package for gRPC, here in this project, package for gRPC is hello_grpc <br />
in this directory create your .proto<br />
in the same directory of your proto file and write this command : <br />
protoc --proto_path=grpc_hello --go_out=grpc_hello hello.proto  //This command creates .pb.go file<br />
protoc --proto_path=grpc_hello --go-grpc_out=grpc_hello hello.proto  //This command creates _grpc.pb.go<br />

Because I want to call the server in the main function I write it as a function and for calling it cause it gets the terminal and does not let other lines that are after calling the server be processed  I use goroutine <br />

this line : option go_package = "./grpchellopbs"; in .proto file indicates the dicrectory of  those .pb and _grpc.pb files and it should be defined otherwise you will face an error 


 
