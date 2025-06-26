package handler

import (
	"time"

	"app/request"
	"app/response"
	"app/service"
	"app/util"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// @Summary 登入
// @Description 登入
// @Tags auth
// @Accept x-www-form-urlencoded
// @Param username formData string true "使用者名稱"
// @Param password formData string true "密碼"
// @Success 200 {object} response.SuccessResponseHTTP{data=response.TokenResponse}
// @Failure 400 {object} response.ErrorResponseHTTP{}
// @Failure 401 {object} response.ErrorResponseHTTP{}
// @Failure 500 {object} response.ErrorResponseHTTP{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {

	input := new(request.Login)

	if err := c.BodyParser(input); err != nil {
		return response.NewErrorRes(fiber.StatusBadRequest, []string{"輸入資料錯誤"})
	}

	lang := c.Get("Accept-Language")

	v := util.NewValidator(lang)

	if err := v.ValidateStruct(input); err != nil {
		return response.NewErrorRes(fiber.StatusBadRequest, err)
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

	return c.JSON(response.NewSuccessRes(response.TokenResponse{
		AccessToken: accessToken,
	}))
}

// @Summary 刷新 access token
// @Description 刷新 access token
// @Tags auth
// @Success 200 {object} response.SuccessResponseHTTP{data=response.TokenResponse}
// @Failure 401 {object} response.ErrorResponseHTTP{}
// @Failure 500 {object} response.ErrorResponseHTTP{}
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if refreshToken == "" {
		return response.NewErrorRes(fiber.StatusUnauthorized, []string{"請提供有效的 refresh token"})
	}

	accessToken, err := h.authService.RefreshToken(refreshToken)

	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessRes(response.TokenResponse{
		AccessToken: accessToken,
	}))
}
