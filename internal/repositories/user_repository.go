package repositories

import (
	"user-service/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx *fiber.Ctx, db *gorm.DB, userModel *models.User) (int64, error)
	CreateUserBio(ctx *fiber.Ctx, db *gorm.DB, userModel *models.User, userId float64) (map[string]interface{}, error)
	FindUserByEmail(ctx *fiber.Ctx, db *gorm.DB, email string) *models.User
	FindUserByPhone(ctx *fiber.Ctx, db *gorm.DB, phone string) *models.User
	FindUserById(ctx *fiber.Ctx, db *gorm.DB, userId float64) (*models.User, error)
}
