package controller

import (
	"db-go-gin/internal/app/dto/request"
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/app/service"
	"db-go-gin/internal/global/statuscode"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/multi"
)

type IUserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	userService service.IUserService
}

func NewUserController() IUserController {
	return &userController{
		userService: service.NewUserService(),
	}
}

// Login 登录
// @Summary 登录
// @Tags API.user
// @Accept application/json
// @Produce application/json
// @Param data body request.LoginReq true "登录参数"
// @Success 200 {object} response.LoginResp "{"code":200,"msg":"ok","data":{}}"
// @Router /user/login [post]
func (u *userController) Login(ctx *gin.Context) {
	var req request.LoginReq
	if !validateJsonData(ctx, &req) {
		return
	}
	resp := u.userService.Login(req)
	response.Response(ctx, resp)
}

// Register 注册账号
// @Summary 注册账号
// @Tags API.user
// @Accept application/json
// @Produce application/json
// @Param data body request.RegisterReq{model.UserBase,model.GatewayBase,model.ApplicationChainBase,model.CrossChainChannelBase} true "注册参数"
// @Success 200 {object} response.EmptyResp "{"code":200,"msg":"ok","data":{}}"
// @Router /user/register [post]
func (u *userController) Register(ctx *gin.Context) {
	var req request.RegisterReq
	if !validateJsonData(ctx, &req) {
		return
	}
	resp := u.userService.Register(req)
	response.Response(ctx, resp)
}

// Logout 登出
// @Summary 登出
// @Tags API.user
// @Security LoginToken
// @Produce application/json
// @Success 200 {object} response.EmptyResp "{"code":200,"msg":"ok","data":{}}"
// @Router /user/logout [post]
func (u *userController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if err := multi.AuthDriver.DelUserTokenCache(token); err != nil {
		response.Response(ctx, response.NewErrorRespMsg(statuscode.SystemError, statuscode.GetText(statuscode.SystemError)))
		return
	}
	response.Response(ctx, response.NewSuccessRespMsg(""))
}
