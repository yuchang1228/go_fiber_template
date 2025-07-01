package middleware

import (
	"app/internal/response"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return c.Status(e.Code).JSON(fiber.Map{
			"success": false,
			"msg":     []string{err.Error()},
		})
	}

	if err != nil {
		if httpErr, ok := err.(*response.HTTPError); ok {
			return c.Status(httpErr.Code).JSON(fiber.Map{
				"success": false,
				"msg":     httpErr.Msg,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     []string{"伺服器錯誤"},
		})
	}
	return nil
}
