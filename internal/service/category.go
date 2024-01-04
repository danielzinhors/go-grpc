package service

import (
	"context"

	"github.com/danielzinhors/go-grpc/internal/database"
	"github.com/danielzinhors/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil

}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesReponse []*pb.Category

	for _, category := range categories {
		categorieReponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesReponse = append(categoriesReponse, categorieReponse)
	}
	return &pb.CategoryList{Categories: categoriesReponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Find(in.Id)
	if err != nil {
		return nil, err
	}

	categorieReponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{Category: categorieReponse}, nil
}
