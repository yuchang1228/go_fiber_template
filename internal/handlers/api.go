package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Health check
// @Summary Health check
// @Description Health check
// @Tags app
// @Success 200 {object} response.SuccessResponseHTTP{data=string}
// @Failure 500 {object} response.ErrorResponseHTTP{}
// @routes /health [get]
func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": "true",
		"data":    "Hello, World!",
	})
}
