package handlers

import (
	"app/internal/requests"
	"app/internal/responses"
	"app/internal/services"
	"app/pkg/validator"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// @Summary 登入
// @Description 登入
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param username formData string true "使用者名稱"
// @Param password formData string true "密碼"
// @Success 200 {object} responses.SuccessResponseHTTP{data=responses.TokenResponse}
// @Failure 400 {object} responses.ErrorResponseHTTP{}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {

	input := new(requests.Login)

	if err := c.BodyParser(input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"輸入資料錯誤"})
	}

	lang := c.Get("Accept-Language")

	v := validator.NewValidator(lang)

	if err := v.ValidateStruct(input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, err)
	}

	accessToken, refreshToken, err := h.authService.Login(input)

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(responses.NewSuccessRes(responses.TokenResponse{
		AccessToken: accessToken,
	}))
}

// @Summary 刷新 access token
// @Description 刷新 access token
// @Tags auth
// @Success 200 {object} responses.SuccessResponseHTTP{data=responses.TokenResponse}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if refreshToken == "" {
		return responses.NewErrorRes(fiber.StatusUnauthorized, []string{"請提供有效的 refresh token"})
	}

	accessToken, err := h.authService.RefreshToken(refreshToken)

	if err != nil {
		return err
	}

	return c.JSON(responses.NewSuccessRes(responses.TokenResponse{
		AccessToken: accessToken,
	}))
}
