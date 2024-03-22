package test

import (
	"dimasfadilah/gin-gorm-clean-arch/controller"
	"dimasfadilah/gin-gorm-clean-arch/middleware"
	"dimasfadilah/gin-gorm-clean-arch/repository"
	"dimasfadilah/gin-gorm-clean-arch/routes"
	"dimasfadilah/gin-gorm-clean-arch/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func closeTestDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func setupRouter() *gin.Engine {
	db := setupTestDB()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	server := gin.Default()
	server.Use(middleware.CORS())

	routes.CategoryRoutes(server, categoryController)

	return server
}

func truncateCategory(db *gorm.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySucces(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	server := setupRouter()
	defer closeTestDB(db)

	// todo
	requestBody := `{"name": "test"}`
	request := httptest.NewRequest("POST", "/category", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	server.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}
