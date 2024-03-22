package helper

import (
	"dimasfadilah/gin-gorm-clean-arch/dto"
	"dimasfadilah/gin-gorm-clean-arch/model"
)

func ToCategoryGeneralResponse(category model.Category) dto.CategoryGeneralResponse {
	return dto.CategoryGeneralResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryGeneralResponses(categories []model.Category) []dto.CategoryGeneralResponse {
	var CategoryGeneralResponses []dto.CategoryGeneralResponse
	for _, category := range categories {
		CategoryGeneralResponses = append(CategoryGeneralResponses, ToCategoryGeneralResponse(category))
	}
	return CategoryGeneralResponses
}
