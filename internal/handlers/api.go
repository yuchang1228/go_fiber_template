package handlers

import (
	"app/internal/tasks"
	"fmt"

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

	_, err := tasks.SendAddTask()
	if err != nil {
		fmt.Print("Could not send task")
	}

	return c.JSON(fiber.Map{
		"success": "true",
		"data":    "Hello, World!",
	})
}
