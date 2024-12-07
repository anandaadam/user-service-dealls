package middleware

import (
	"os"
	"strings"
	"time"
	"user-service/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return response.ErrorResponse(ctx, "Missing Authorization header", fiber.StatusUnauthorized)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return response.ErrorResponse(ctx, "Invalid Authorization format", fiber.StatusUnauthorized)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			return response.ErrorResponse(ctx, "Invalid or expired token", fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return response.ErrorResponse(ctx, "Invalid token", fiber.StatusUnauthorized)
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return response.ErrorResponse(ctx, "Token is expired", fiber.StatusUnauthorized)
			}
		} else {
			return response.ErrorResponse(ctx, "Expiration claim missing", fiber.StatusUnauthorized)
		}

		ctx.Locals("jwt", token)
		return ctx.Next()
	}
}
