package service

import (
	"app/repository"
	"app/request"
	"app/response"
	"app/util"

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
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{util.GormErrorToMessage(err)})
	}

	if !util.CheckPasswordHash(input.Password, user.Password) {
		return "", "", response.NewErrorRes(fiber.StatusUnauthorized, []string{"使用者名稱或密碼錯誤"})
	}

	accessToken, err := util.GenerateAccessJWT(user.ID, user.Username, user.Email)

	if err != nil {
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"access token 生成失敗"})
	}

	refreshTokoen, err := util.GenerateRefreshJWT(user.ID, user.Username)

	if err != nil {
		return "", "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"refresh token 生成失敗"})
	}

	return accessToken, refreshTokoen, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, *response.HTTPError) {
	userID, err := util.ParseRefreshJWT(refreshToken)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusUnauthorized, []string{"refresh token 解析失敗"})
	}

	user, err := s.userRepository.GetByID(userID)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + util.GormErrorToMessage(err)})
	}

	accessToken, err := util.GenerateAccessJWT(user.ID, user.Username, user.Email)

	if err != nil {
		return "", response.NewErrorRes(fiber.StatusInternalServerError, []string{"access token 生成失敗"})
	}

	return accessToken, nil
}
