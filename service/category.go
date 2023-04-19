package service

import (
	"context"

	"github.com/felipedias-dev/fullcycle-go-expert-grpc/internal/database"
	"github.com/felipedias-dev/fullcycle-go-expert-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	category, err := cs.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
