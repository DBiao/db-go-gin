package controller

import (
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/global/statuscode"

	"github.com/gin-gonic/gin"
)

// validateJsonData bing data to Json
func validateJsonData(ctx *gin.Context, model interface{}) bool {
	err := ctx.ShouldBindJSON(&model)
	if err != nil {
		response.Response(ctx, response.NewErrorRespMsg(statuscode.SystemParamsError, statuscode.GetText(statuscode.SystemParamsError)))
		ctx.Abort()
		return false
	}
	return true
}
