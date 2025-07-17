package requests

type Login struct {
	// 使用者名稱
	Username string `json:"username" validate:"required,min=3,max=20"`

	// 密碼
	Password string `json:"password" validate:"required,min=6"`
}
