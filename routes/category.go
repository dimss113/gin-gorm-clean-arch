package routes

import (
	"dimasfadilah/gin-gorm-clean-arch/controller"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, categoryController controller.CategoryController) {
	router.GET("/categories", categoryController.FindAllCategory)
	router.GET("/category/:id", categoryController.FindCategoryByID)
	router.POST("/category", categoryController.CreateCategory)
}
