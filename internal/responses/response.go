package responses

import "fmt"

/*
HTTP Success Response
*/
type SuccessResponseHTTP struct {
	// 狀態
	Success bool `json:"success"`

	Data any `json:"data"`
} // @name SuccessResponseHTTP

func NewSuccessRes(data any) *SuccessResponseHTTP {
	return &SuccessResponseHTTP{
		Success: true,
		Data:    data,
	}
}

/*
HTTP Error Reponse
*/
type ErrorResponseHTTP struct {
	// 狀態
	Success bool `json:"success" example:"false"`

	// 錯誤訊息陣列
	Msg []string `json:"msg"`
} // @name ErrorResponseHTTP

type HTTPError struct {
	Code int
	Msg  []string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP Error %d: %v", e.Code, e.Msg)
}

func NewErrorRes(code int, msg []string) *HTTPError {
	return &HTTPError{
		Code: code,
		Msg:  msg,
	}
}
