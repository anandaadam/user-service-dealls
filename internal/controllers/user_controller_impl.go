package controllers

import (
	"user-service/internal/models"
	"user-service/internal/services"
	"user-service/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	reqBody := &models.CreateUserRequest{}

	if err := ctx.BodyParser(reqBody); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(reqBody)

	if err != nil {
		errorMessages := errorMessageValidation(err.(validator.ValidationErrors))
		return response.FailResponse(ctx, "Error validation", errorMessages, 422)
	}

	user, err := userController.UserService.CreateUser(ctx, reqBody)

	if err != nil {
		return response.FailResponse(ctx, "Failed to create user", err.Error(), 500)
	}

	return response.SuccessResponse(ctx, "Success to signup", user, 201)
}

func (userController *UserControllerImpl) CreateUserBio(ctx *fiber.Ctx) error {
	reqBody := &models.CreateUserBioRequest{}

	if err := ctx.BodyParser(reqBody); err != nil {
		return err
	}

	user, err := userController.UserService.CreateUserBio(ctx, reqBody)

	if err != nil {
		return response.FailResponse(ctx, "Failed to create user bio", err.Error(), 500)
	}

	return response.SuccessResponse(ctx, "Success to create user bio", user, 201)
}

func (userController *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	reqBody := &models.UserLoginRequest{}

	if err := ctx.BodyParser(reqBody); err != nil {
		return err
	}

	user, err := userController.UserService.Login(ctx, reqBody)

	if err != nil {
		return response.FailResponse(ctx, "Failed to login", err.Error(), 500)
	}

	return response.SuccessResponse(ctx, "Success to login", user, 200)
}

func errorMessageValidation(errors validator.ValidationErrors) []string {
	messages := make([]string, 0)

	for _, err := range errors {
		switch err.Field() {
		case "Email":
			switch err.Tag() {
			case "required":
				messages = append(messages, "Email wajib diisi")
			case "email":
				messages = append(messages, "Email format salah")
			}

		case "Password":
			switch err.Tag() {
			case "required":
				messages = append(messages, "Password wajib diisi")
			case "min":
				messages = append(messages, "Password minimal 8 karakter")
			}
		}
	}

	return messages
}
