package statuscode

const (
	SUCCESS = 200
	ERROR   = 500

	// SystemError 系统错误
	SystemError           = 100000 // 系统错误
	SystemPermissionError = 100001 // 权限不足
	SystemUserAuthError   = 100002 // 用户认证错误
	SystemUploadImgError  = 100003 // 上传图片错误
	SystemParamsError     = 100004 // 参数错误

	UserLoginAccountError    = 100100 // 用户名错误
	UserLoginPasswordError   = 100101 // 密码错误
	RegisterAccountTypeError = 100102 // 注册账号类型错误
	RegisterAccountError     = 100103 // 注册账号错误
	AccountNoAuditError      = 100104 // 注册审核不通过
	UpdateUserPasswordError  = 100105 // 修改密码错误
	UpdateUserInfoError      = 100106 // 修改用户信息错误

)

// GetText 获取错误码文本
func GetText(code int) string {
	return zhCNText[code]
}
