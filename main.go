package main

import (
	gen "kevin/gogrpc/pb"
	"kevin/gogrpc/repo"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	//create new grpc server
	//register service
	//run server
	grpc_server := grpc.NewServer()
	gen.RegisterProductServiceServer(grpc_server, repo.NewProductImpl())
	gen.RegisterTodoServiceServer(grpc_server, repo.NewTodoImpl())
	PORT := ":9000"
	if l, err := net.Listen("tcp", PORT); err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	} else {
		// the gRPC server
		if err := grpc_server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}

	}

}
