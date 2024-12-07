package services

import (
	"user-service/internal/models"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx, userRequest *models.CreateUserRequest) (string, error)
	CreateUserBio(ctx *fiber.Ctx, userRequest *models.CreateUserBioRequest) (map[string]interface{}, error)
	Login(ctx *fiber.Ctx, userRequest *models.UserLoginRequest) (string, error)
}
