package request

import "db-go-gin/internal/app/model"

type LoginReq struct {
	Account  string `json:"account"`  // 账号
	Password string `json:"password"` // 密码
}

type RegisterReq struct {
	AccountType uint8 `json:"account_type"` // 账号类型 1:网关管理员 2.应用链管理员 3.跨链通道管理员
	model.UserBase
}
