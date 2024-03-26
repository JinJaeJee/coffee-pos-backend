package main

import (
	"coffee-pos-backend/middlewares"
	"coffee-pos-backend/models"
	"coffee-pos-backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initDB() {
	dsn := "host=localhost user=youruser dbname=yourdb password=yourpassword port=5430 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected")

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.RolePermission{})

}

func main() {
	initDB()
	r := gin.Default()
	r.Use(middlewares.VerifyToken())

	// SeedRoles(db)
	routes.UserRoutes(r, db)
	routes.AuthRoutes(r, db)
	r.GET("/checkapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to POS-Coffee-Cafe API!",
		})
	})

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server")
	}

}
