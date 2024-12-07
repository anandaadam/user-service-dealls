package main

import (
	"log"
	"user-service/database"
	"user-service/internal/controllers"
	"user-service/internal/repositories"
	"user-service/internal/services"
	"user-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	db, _ := database.GetDBConnection()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	router := routes.Router(userController)

	router.Listen(":3000")
}
