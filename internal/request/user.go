package request

type CreateUser struct {
	// 使用者名稱
	Username string `json:"username" validate:"required,min=3,max=20"`

	// 電子郵件
	Email string `json:"email" validate:"required,email"`

	// 密碼
	Password string `json:"password" validate:"required,min=6,max=50"`

	// 姓名
	Names string `json:"names"`
}

type UpdateUser struct {
	// 姓名
	Names string `json:"names" validate:"required,min=1,max=50"`
}
