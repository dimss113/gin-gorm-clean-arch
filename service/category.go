package service

import (
	"context"
	"dimasfadilah/gin-gorm-clean-arch/dto"
	"dimasfadilah/gin-gorm-clean-arch/model"
	"dimasfadilah/gin-gorm-clean-arch/repository"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, req dto.CategoryUploadRequest) (dto.CategoryGeneralResponse, error)
	FindCategoryByID(ctx context.Context, id uint64) (dto.CategoryGeneralResponse, error)
	FindAllCategory(ctx context.Context) ([]dto.CategoryGeneralResponse, error)
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryServiceImpl) CreateCategory(ctx context.Context, req dto.CategoryUploadRequest) (dto.CategoryGeneralResponse, error) {
	// create category
	category := model.Category{
		Name: req.Name,
	}

	result, err := c.categoryRepository.CreateCategory(ctx, nil, category)
	if err != nil {
		return dto.CategoryGeneralResponse{}, err
	}

	return dto.CategoryGeneralResponse{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}

func (c *CategoryServiceImpl) FindCategoryByID(ctx context.Context, id uint64) (dto.CategoryGeneralResponse, error) {
	result, err := c.categoryRepository.FindOneByID(ctx, nil, id)
	if err != nil {
		return dto.CategoryGeneralResponse{}, err
	}

	return dto.CategoryGeneralResponse{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}

func (c *CategoryServiceImpl) FindAllCategory(ctx context.Context) ([]dto.CategoryGeneralResponse, error) {
	result, err := c.categoryRepository.FindAll(ctx, nil)
	if err != nil {
		return []dto.CategoryGeneralResponse{}, err
	}

	var response []dto.CategoryGeneralResponse
	for _, category := range result {
		response = append(response, dto.CategoryGeneralResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return response, nil
}
