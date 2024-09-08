package main

import (
	"code/controllers"
	"code/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	wd, _ := os.Getwd()

	curDir := fmt.Sprint(wd, "/.env")

	err = godotenv.Load(curDir)

	if err != nil {
		log.Fatal("INFRASTRUCTURE:1001): ", err)
	}

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"),
		os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", sqlInfo)
	defer func() {
		if err := DB.Close(); err != nil {
			fmt.Println("Error : ", err)
		}
	}()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB, "up")

	fmt.Println("Successfully connected!")

	router := gin.Default()
	router.POST("/api/register", controllers.RegisterUser)
	router.POST("/api/login", controllers.LoginUser)
	router.GET("/api/categories", controllers.GetAllCategories)
	router.GET("/api/categories/:id", controllers.GetCategoryDetails)
	router.POST("/api/categories", controllers.InsertCategory)
	router.POST("/api/books", controllers.InsertBooks)
	router.DELETE("/api/categories/:id", controllers.DeleteCat)
	router.GET("/api/books", controllers.GetAllBooks)
	router.GET("/api/books/:id", controllers.GetBookDet)
	router.GET("/api/categories/:id/books", controllers.GetCategoryBooks)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
