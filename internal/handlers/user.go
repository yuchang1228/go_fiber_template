package handlers

import (
	"app/internal/models"
	"app/internal/requests"
	"app/internal/responses"
	"app/internal/services"
	"app/pkg/bcrypt"
	"app/pkg/gorm"
	"app/pkg/validator"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService       services.IUserService
	userReportService services.IUserReportService
}

func NewUserHandler(userService services.IUserService, userReportService services.IUserReportService) *UserHandler {
	return &UserHandler{userService, userReportService}
}

type CreateUserRequest struct {
	// 使用者名稱
	Username string `json:"username" validate:"required"`

	// 電子郵件
	Email string `json:"email" validate:"required"`

	// 密碼
	Password string `json:"password" validate:"required"`

	// 姓名
	Names string `json:"names"`
}

type UpdateUserRequest struct {
	// 姓名
	Names string `json:"names" validate:"required"`
}

type User struct {
	// 使用者名稱
	Username string `json:"username"`

	// 電子郵件
	Email string `json:"email"`

	// 密碼
	Password string `json:"password"`

	// 姓名
	Names string `json:"names"`
}

// @Summary 取得所有使用者
// @Description 取得所有使用者
// @Tags user
// @Accept json
// @Success 200 {object} responses.SuccessResponseHTTP{data=[]responses.UserResponse}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user [get]
// @Security Bearer
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)},
		})
	}

	var result []responses.UserResponse

	for _, user := range *users {
		result = append(result, responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Names:     user.Names,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessRes(result))
}

// @Summary 取得使用者透過ID
// @Description 取得使用者透過ID
// @Tags user
// @Param id path string true "ID"
// @Success 200 {object} responses.SuccessResponseHTTP{data=responses.UserResponse}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user/{id} [get]
// @Security Bearer
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	idUint, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"URL 參數格式錯誤"})
	}

	user, err := h.userService.GetByID(uint(idUint))

	if err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	return c.JSON(responses.NewSuccessRes(responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Names:     user.Names,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}))
}

// @Summary 新增使用者
// @Description 新增使用者
// @Tags user
// @Accept x-www-form-urlencoded
// @Param username formData string true "使用者名稱"
// @Param email formData string true "電子郵件"
// @Param password formData string true "密碼"
// @Param names formData string false "姓名"
// @Success 200 {object} responses.SuccessResponseHTTP{data=responses.UserResponse}
// @Success 400 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user [post]
// @Security Bearer
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	input := new(requests.CreateUser)
	if err := c.BodyParser(input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"資料格式錯誤"})
	}

	lang := c.Get("Accept-Language")

	v := validator.NewValidator(lang)

	if err := v.ValidateStruct(input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, err)
	}

	hash, err := bcrypt.HashPassword(input.Password)

	if err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"密碼加密失敗"})
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
		Names:    input.Names,
	}

	if err := h.userService.Create(&user); err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	return c.JSON(responses.NewSuccessRes(responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Names:     user.Names,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}))
}

// @Summary 編輯使用者
// @Description 編輯使用者
// @Tags user
// @Accept x-www-form-urlencoded
// @Param id path string true "ID"
// @Param names formData string false "姓名"
// @Success 200 {object} responses.SuccessResponseHTTP{data=responses.UserResponse}
// @Success 400 {object} responses.ErrorResponseHTTP{}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user/{id} [patch]
// @Security Bearer
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var input requests.UpdateUser

	id := c.Params("id")

	idUint, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"URL 參數格式錯誤"})
	}

	if err := c.BodyParser(&input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"資料格式錯誤"})
	}

	lang := c.Get("Accept-Language")

	v := validator.NewValidator(lang)

	if err := v.ValidateStruct(input); err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, err)
	}

	user, err := h.userService.GetByID(uint(idUint))

	if err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	user.Names = input.Names

	err = h.userService.Update(user)

	if err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	return c.JSON(responses.NewSuccessRes(user))
}

// @Summary 刪除使用者
// @Description 刪除使用者
// @Tags user
// @Param id path string true "User ID"
// @Success 200 {object} responses.SuccessResponseHTTP{data=nil}
// @Success 400 {object} responses.ErrorResponseHTTP{}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user/{id} [delete]
// @Security Bearer
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	idUint, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return responses.NewErrorRes(fiber.StatusBadRequest, []string{"URL 參數格式錯誤"})
	}

	if err := h.userService.Delete(uint(idUint)); err != nil {
		return responses.NewErrorRes(fiber.StatusInternalServerError, []string{"資料庫錯誤: " + gorm.GormErrorToMessage(err)})
	}

	return c.JSON(responses.NewSuccessRes(nil))
}

// @Summary 匯出使用者報表
// @Description 匯出使用者報表
// @Tags user
// @Success 200 {object} responses.SuccessResponseHTTP{data=nil}
// @Failure 401 {object} responses.ErrorResponseHTTP{}
// @Failure 500 {object} responses.ErrorResponseHTTP{}
// @Router /user/report [get]
// @Security Bearer
func (h *UserHandler) UserReport(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()

	if err != nil {
		return err
	}

	data, err := h.userReportService.GenerateExcel(users)

	if err != nil {
		return err
	}

	filename := "用戶清單.xlsx"
	encodedFilename := url.QueryEscape(filename)
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", `attachment; filename="fallback.xlsx"; filename*=UTF-8''`+encodedFilename)
	return c.Send(data)
}
