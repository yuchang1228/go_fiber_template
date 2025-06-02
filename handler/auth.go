package handler

import (
	"time"

	"app/service"
	"app/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService service.IUserService
}

func NewAuthHandler(userService service.IUserService) *AuthHandler {
	return &AuthHandler{userService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Password string `json:"password" validate:"required,min=6"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "輸入資料錯誤"})
	}

	if err := util.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.TranslateErrors(err.(validator.ValidationErrors), map[string]string{
			"Username": "使用者名稱",
			"Password": "密碼",
		}))
	}

	username := input.Username
	password := input.Password

	user, err := h.userService.GetUserByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "資料庫錯誤", "errors": err.Error()})
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "密碼錯誤"})
	}

	accessToken, err := util.GenerateAccessJWT(user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refreshTokoen, err := util.GenerateRefreshJWT(user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokoen,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{"success": true, "message": "登入成功", "data": fiber.Map{
		"access_token": accessToken,
	}})
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	return c.SendString("test")
}
