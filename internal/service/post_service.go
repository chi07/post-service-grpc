package service

import (
	"context"
	"log"
	"strconv"

	"github.com/chi07/proto/proto-gen-go/pb" // Use this from external package

	"post-service-grpc/internal/repository"
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
