package main

import (
	"api-integration/config"
	"api-integration/controllers"
	"api-integration/middleware"
	"api-integration/models"
	"api-integration/rabbitmq"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	db := config.ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Tags{}, &models.PostTag{})

	authController := controllers.NewAuthController(db)
	postController := controllers.NewPostController(db)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}

	// Set up RabbitMQ connection
	rabbitMQService, err := rabbitmq.NewRabbitMQService("amqp://user:password@rabbitmq:5672/", "product_queue")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rabbitMQService.Close()

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controllers.NewUserController(db, rabbitMQService).GetUsers)
		protected.POST("/users", controllers.NewUserController(db, rabbitMQService).CreateUser)

		// Without DB
		protected.POST("/users_without_db", controllers.CreateUserWithoutDB)
		protected.GET("/users_without_db", controllers.GetUserWithoutDB)

		// Tag routes
		protected.POST("/tags", postController.CreateTag)
		protected.POST("/posts", postController.CreatePost)
		protected.GET("/posts", postController.GetPosts)
		protected.GET("/posts/:id", postController.GetPostsByID)
	}

	r.Run(":8080")
}
