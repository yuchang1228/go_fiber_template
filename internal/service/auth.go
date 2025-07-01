package service

import (
	"app/internal/repository"
	"app/internal/request"
	"app/internal/response"
	"app/pkg/bcrypt"
	"app/pkg/gorm"
	"app/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type IAuthService interface {
	Login(input *request.Login) (string, string, *response.HTTPError)
	RefreshToken(refreshToken string) (string, *response.HTTPError)
}

type authService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(
	userRepository repository.IUserRepository,
) IAuthService {
	return &authService{userRepository}
}

func (s *authService) Login(input *request.Login) (string, string, *response.HTTPError) {
	user, err := s.userRepository.GetByUserName(input.Username)

	if err != nil {
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{gorm.GormErrorToMessage(err)})
	}

	if !bcrypt.CheckPasswordHash(input.Password, user.Password) {
		return "", "", response.NewErrorRes(fiber.StatusUnauthorized, []string{"使用者名稱或密碼錯誤"})
	}

	accessToken, err := jwt.GenerateAccessJWT(user.ID, user.Username, user.Email)

	if err != nil {
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"access token 生成失敗"})
	}

	refreshTokoen, err := jwt.GenerateRefreshJWT(user.ID, user.Username)

	if err != nil {
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"refresh token 生成失敗"})
	}

	return accessToken, refreshTokoen, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, *response.HTTPError) {
	userID, err := jwt.ParseRefreshJWT(refreshToken)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusUnauthorized, []string{"refresh token 解析失敗"})
	}

	user, err := s.userRepository.GetByID(userID)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	accessToken, err := jwt.GenerateAccessJWT(user.ID, user.Username, user.Email)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"access token 生成失敗"})
	}

	return accessToken, nil
}
