package repositories

import (
	"user-service/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateUser(ctx *fiber.Ctx, tx *gorm.DB, userModel *models.User) (int64, error) {
	err := tx.Create(userModel).Error

	return userModel.Id, err
}

func (repository *UserRepositoryImpl) CreateUserBio(ctx *fiber.Ctx, tx *gorm.DB, userModel *models.User, userId float64) (map[string]interface{}, error) {
	err := tx.Model(userModel).Where("id = ?", userId).Updates(userModel).Error

	return map[string]interface{}{
		"fullname":  userModel.FullName,
		"phone":     userModel.Phone,
		"birthdate": userModel.BirthDate,
		"city":      userModel.City,
		"bio":       userModel.Bio,
	}, err
}

func (repository *UserRepositoryImpl) FindUserByEmail(ctx *fiber.Ctx, db *gorm.DB, email string) *models.User {
	var user models.User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil
	}

	return &user
}

func (repository *UserRepositoryImpl) FindUserByPhone(ctx *fiber.Ctx, db *gorm.DB, phone string) *models.User {
	var user models.User

	result := db.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil
	}

	return &user
}

func (repository *UserRepositoryImpl) FindUserById(ctx *fiber.Ctx, db *gorm.DB, userId float64) (*models.User, error) {
	var user models.User

	result := db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
