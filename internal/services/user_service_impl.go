package services

import (
	"encoding/json"
	"fmt"
	"time"
	"user-service/config"
	"user-service/internal/models"
	"user-service/internal/repositories"
	jwttoken "user-service/pkg/jwt_token"
	"user-service/pkg/password"
	"user-service/pkg/username"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repositories.UserRepository, db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (userService *UserServiceImpl) CreateUser(ctx *fiber.Ctx, userRequest *models.CreateUserRequest) (string, error) {
	existingUser := userService.UserRepository.FindUserByEmail(ctx, userService.DB, userRequest.Email)
	if existingUser != nil {
		return "", fmt.Errorf("email '%s' sudah terdaftar", userRequest.Email)
	}

	password, err := password.HashPassword(userRequest.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	userRequest.Password = password

	username := username.GenerateUsername(userRequest.Email)

	userModel := &models.User{}
	userModel.Email = userRequest.Email
	userModel.Password = userRequest.Password
	userModel.Username = username

	var userId int64
	err = userService.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		userId, err = userService.UserRepository.CreateUser(ctx, tx, userModel)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	var token string
	token, err = jwttoken.GenerateToken(userId)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (userService *UserServiceImpl) CreateUserBio(ctx *fiber.Ctx, userRequest *models.CreateUserBioRequest) (map[string]interface{}, error) {
	existingUser := userService.UserRepository.FindUserByPhone(ctx, userService.DB, userRequest.Phone)
	if existingUser != nil {
		return map[string]interface{}{}, fmt.Errorf("telepon '%s' sudah terdaftar", userRequest.Phone)
	}

	token := ctx.Locals("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	birthDate, err := time.Parse("2006-01-02", userRequest.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("invalid birthdate format: %w", err)
	}

	userModel := &models.User{}
	userModel.FullName = userRequest.FullName
	userModel.Phone = &userRequest.Phone
	userModel.BirthDate = birthDate
	userModel.City = userRequest.City
	userModel.Bio = userRequest.Bio

	var userBio map[string]interface{}
	err = userService.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		userBio, err = userService.UserRepository.CreateUserBio(ctx, tx, userModel, userId)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	user, err := userService.UserRepository.FindUserById(ctx, userService.DB, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	_, err = userService.PublishUserEvent(user)
	if err != nil {
		return nil, fmt.Errorf("publis event failed: %w", err)
	}

	return userBio, nil
}

func (userService *UserServiceImpl) Login(ctx *fiber.Ctx, userRequest *models.UserLoginRequest) (string, error) {
	existingUser := userService.UserRepository.FindUserByEmail(ctx, userService.DB, userRequest.Email)
	if existingUser == nil {
		return "", fmt.Errorf("email '%s' tidak terdaftar", userRequest.Email)
	}

	if !password.CheckPasswordHash(userRequest.Password, existingUser.Password) {
		return "", fmt.Errorf("password salah")
	}

	token, err := jwttoken.GenerateToken(existingUser.Id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (userService *UserServiceImpl) PublishUserEvent(user *models.User) (string, error) {
	producer, err := kafka.NewProducer(config.KafkaConfig())
	if err != nil {
		return "", fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	defer producer.Close()

	messageBody := map[string]interface{}{
		"eventType": "create_user",
		"data":      user,
	}

	messageValue, err := json.Marshal(messageBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal Kafka message body: %w", err)
	}

	topic := "user"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: messageValue,
	}, nil)

	if err != nil {
		return "", fmt.Errorf("failed to produce Kafka message: %w", err)
	}

	return "success", nil
}
