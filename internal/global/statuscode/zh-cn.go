package statuscode

// 错误码中文翻译
var zhCNText = map[int]string{
	SystemError:           "系统错误",
	SystemPermissionError: "权限不足",
	SystemUserAuthError:   "用户认证错误",
	SystemUploadImgError:  "上传图片错误",
	SystemParamsError:     "参数错误",

	UserLoginAccountError:  "账号错误",
	UserLoginPasswordError: "密码错误",

	RegisterAccountTypeError: "注册账号类型错误",
	RegisterAccountError:     "注册账号错误",
}
