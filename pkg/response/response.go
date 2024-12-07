package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *fiber.Ctx, message string, data interface{}, statusCode int16) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	return ctx.Status(int(statusCode)).JSON(response)
}

func FailResponse(ctx *fiber.Ctx, message string, data interface{}, statusCode int16) error {
	response := Response{
		Success: false,
		Message: message,
		Data:    data,
	}

	return ctx.Status(int(statusCode)).JSON(response)
}

func ErrorResponse(ctx *fiber.Ctx, message string, statusCode int16) error {
	response := Response{
		Success: false,
		Message: message,
	}

	return ctx.Status(int(statusCode)).JSON(response)
}
