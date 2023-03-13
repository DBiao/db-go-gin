package controller

import (
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/app/service"
	"github.com/gin-gonic/gin"
)

// CaptchaMath 生成验证码
func CaptchaMath(ctx *gin.Context) {
	resp := service.CaptchaMath()
	response.Response(ctx, resp)
}
