package main

import (
	"dimasfadilah/gin-gorm-clean-arch/app"
	"dimasfadilah/gin-gorm-clean-arch/controller"
	"dimasfadilah/gin-gorm-clean-arch/middleware"
	"dimasfadilah/gin-gorm-clean-arch/migration"
	"dimasfadilah/gin-gorm-clean-arch/repository"
	"dimasfadilah/gin-gorm-clean-arch/routes"
	"dimasfadilah/gin-gorm-clean-arch/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	db := app.NewDBConnection()

	// todo
	app.TruncateTable(db, "categories")
	migration.RunMigration(db)
	migration.SeederData(db)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	defer app.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.CategoryRoutes(server, categoryController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
