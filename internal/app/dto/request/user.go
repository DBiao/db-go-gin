package request

import "db-go-gin/internal/app/model"

type LoginReq struct {
	Account  string `json:"account"`  // 账号
	Password string `json:"password"` // 密码
}

type RegisterReq struct {
	model.UserBase
}
