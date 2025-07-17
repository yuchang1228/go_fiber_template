package responses

import "time"

type UserResponse struct {
	// ID
	ID uint `json:"id"`

	// 電子郵件
	Email string `json:"email"`

	// 使用者名稱
	Username string `json:"username"`

	// 姓名
	Names string `json:"names"`

	// 創建時間
	CreatedAt time.Time `json:"createdAt"`

	// 更新時間
	UpdatedAt time.Time `json:"updatedAt"`
} //@name UserResponse
