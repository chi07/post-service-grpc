package main

import (
	"log"
	"net"

	"github.com/chi07/proto/proto-gen-go/pb" // Use this from external package

	"google.golang.org/grpc"

	"post-service-grpc/internal/service"
	"post-service-grpc/pkg/db"
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
