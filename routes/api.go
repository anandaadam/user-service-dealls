package routes

import (
	"user-service/internal/controllers"
	"user-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(userController controllers.UserController) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")

	signup := api.Group("/signup")
	signup.Post("/user", userController.CreateUser)

	login := api.Group("/login")
	login.Post("/user", userController.Login)

	user := api.Group("/user")
	user.Use(middleware.JWTProtected())
	user.Post("/bio", userController.CreateUserBio)

	return app
}
