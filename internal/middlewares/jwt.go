package middlewares

import (
	"app/config"
	"app/internal/responses"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Config("ACCESS_JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.NewErrorRes(fiber.StatusUnauthorized, []string{"JWT 驗證失敗"})
		},
	})
}
