package main

import (
	"log"
	"net"

	"post-service-grpc/internal/service"
	"post-service-grpc/pkg/db"
	pb "post-service-grpc/proto-gen-go/proto"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, &service.PostService{})

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
