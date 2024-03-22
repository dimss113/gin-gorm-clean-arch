package controller

import (
	"dimasfadilah/gin-gorm-clean-arch/dto"
	"dimasfadilah/gin-gorm-clean-arch/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	FindCategoryByID(ctx *gin.Context)
	FindAllCategory(ctx *gin.Context)
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) CreateCategory(ctx *gin.Context) {
	var categoryDTO dto.CategoryUploadRequest
	errDTO := ctx.ShouldBind(&categoryDTO)
	if errDTO != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Status:  dto.STATUS_ERROR,
			Message: errDTO.Error(),
		})
		return
	}

	result, err := c.categoryService.CreateCategory(ctx, categoryDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.STATUS_ERROR,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: dto.MESSAGE_CATEGORY_CREATE_SUCCESS,
		Status:  dto.STATUS_SUCCESS,
		Data:    result,
	})
}

func (c *CategoryControllerImpl) FindCategoryByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Status:  dto.STATUS_ERROR,
			Message: err.Error(),
		})
		return
	}

	result, err := c.categoryService.FindCategoryByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.STATUS_ERROR,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: dto.MESSAGE_CATEGORY_FETCH_SUCCESS,
		Status:  dto.STATUS_SUCCESS,
		Data:    result,
	})
}

func (c *CategoryControllerImpl) FindAllCategory(ctx *gin.Context) {
	result, err := c.categoryService.FindAllCategory(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.STATUS_ERROR,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: dto.MESSAGE_CATEGORY_FETCH_SUCCESS,
		Status:  dto.STATUS_SUCCESS,
		Data:    result,
	})
}
