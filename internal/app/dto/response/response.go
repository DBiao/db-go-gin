package response

import (
	"db-go-gin/internal/global/statuscode"

	"github.com/gin-gonic/gin"
)

type ResponseMsg struct {
	Code    int         `json:"code"`    // 返回码
	Message string      `json:"message"` // 返回信息
	Data    interface{} `json:"data"`    // 返回数据
}

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

type EmptyResp struct {
}

func NewErrorRespMsg(code int, msg string) ResponseMsg {
	return ResponseMsg{Code: code, Message: msg, Data: ""}
}

func NewSuccessRespMsg(data interface{}) ResponseMsg {
	return ResponseMsg{Code: statuscode.SUCCESS, Message: "ok", Data: data}
}

func NewResponseMsg(code int, msg string, data interface{}) ResponseMsg {
	return ResponseMsg{Code: code, Message: msg, Data: data}
}

func Success(ctx *gin.Context, data interface{}, msg ...string) {
	message := "ok"
	if len(msg) > 0 {
		message = msg[0]
	}
	ctx.JSON(statuscode.SUCCESS, NewResponseMsg(statuscode.SUCCESS, message, data))
}

func Response(ctx *gin.Context, resp ResponseMsg) {
	ctx.JSON(statuscode.SUCCESS, resp)
}
