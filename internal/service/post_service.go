package service

import (
	"context"
	"log"
	"strconv"

	"post-service-grpc/internal/repository"
	pb "post-service-grpc/proto-gen-go/proto"
)

// PostService struct
type PostService struct {
	pb.UnimplementedPostServiceServer
}

// CreateArticle handles article creation
func (s *PostService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	id, err := repository.CreateArticle(req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	log.Printf("Article created with ID: %d", id)
	return &pb.CreateArticleResponse{Id: strconv.Itoa(int(id))}, nil
}

// EditArticle handles article editing
func (s *PostService) EditArticle(ctx context.Context, req *pb.EditArticleRequest) (*pb.EditArticleResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = repository.EditArticle(uint(id), req.Title, req.Content)
	if err != nil {
		return &pb.EditArticleResponse{Success: false}, nil
	}
	log.Printf("Article with ID %d updated", id)
	return &pb.EditArticleResponse{Success: true}, nil
}
