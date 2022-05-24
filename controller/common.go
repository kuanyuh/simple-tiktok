package controller


const (
	USERNAME_INVALID = 2 //用户名无效
	PASSWORD_INVALID = 3 //密码无效
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
