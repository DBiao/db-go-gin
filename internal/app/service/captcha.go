package service

import (
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/global"
	"db-go-gin/internal/global/statuscode"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// Math 配置参数
var (
	Height          = 70
	Width           = 240
	NoiseCount      = 0
	ShowLineOptions = base64Captcha.OptionShowHollowLine
	BgColor         = &color.RGBA{
		R: 144,
		G: 238,
		B: 144,
		A: 10,
	}
	FontsStorage base64Captcha.FontsStorage
	Fonts        []string
)

// CaptchaMath 生成验证码
func CaptchaMath() response.ResponseMsg {
	// 这里用的是Math类型的验证码
	driver := base64Captcha.NewDriverMath(Height, Width, NoiseCount, ShowLineOptions, BgColor, FontsStorage, Fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		global.LOG.Error("generate math captcha error", zap.Error(err))
		return response.NewErrorRespMsg(statuscode.SystemError, statuscode.GetText(statuscode.SystemError))
	}

	data := make(map[string]interface{})
	data["id"] = id
	data["b64s"] = b64s

	// 校验验证码
	//if store.Verify(l.CaptchaId, l.Captcha, true) {
	//
	//}else {
	//
	//}

	return response.NewSuccessRespMsg(data)
}
