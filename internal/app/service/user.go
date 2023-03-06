package service

import (
	"db-go-gin/internal/app/dao"
	"db-go-gin/internal/app/dto/request"
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/app/model"
	"db-go-gin/internal/global"
	"db-go-gin/internal/global/statuscode"
	"db-go-gin/internal/utils"
	"strconv"
	"time"

	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

type IUserService interface {
	Login(req request.LoginReq) response.ResponseMsg
	Register(req request.RegisterReq) response.ResponseMsg
}

const (
	NoneAuthority     uint8 = iota // 空授权
	Admin                          // 云平台管理员
	Gateway                        // 网关用户
	ApplicationChain               // 应用链用户
	CrossChainChannel              // 跨链通道用户
)

const (
	NoneStatus  uint8 = iota // 空授权
	Auditing                 // 审核中
	AuditPass                // 审核通过
	AuditRefuse              // 审核拒绝
)

func NewUserService() IUserService {
	return &userService{
		userDao: dao.NewUserDao(),
	}
}

type userService struct {
	userDao dao.IUserDao
}

func (u *userService) Login(req request.LoginReq) response.ResponseMsg {
	// 根据账号获取玩家信息
	user, err := u.userDao.GetByUserName(req.Account)
	if err != nil {
		global.LOG.Error("user account error", zap.String("account", req.Account))
		return response.NewErrorRespMsg(statuscode.UserLoginAccountError, statuscode.GetText(statuscode.UserLoginAccountError))
	}

	// 验证密码
	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		global.LOG.Error("user password error", zap.String("req.Password", utils.BcryptHash(req.Password)), zap.String("user.Password", user.Password))
		return response.NewErrorRespMsg(statuscode.UserLoginPasswordError, statuscode.GetText(statuscode.UserLoginPasswordError))
	}

	// token的信息
	claims := &multi.MultiClaims{
		Id:            strconv.FormatUint(uint64(user.ID), 10),
		Username:      req.Account,
		AuthorityId:   strconv.Itoa(1),
		AuthorityType: int(1),
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresAt:     time.Now().Local().Add(multi.RedisSessionTimeoutWeb).Unix(),
	}

	// 生成token
	token, _, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		global.LOG.Error("generate token error", zap.Error(err))
		return response.NewErrorRespMsg(statuscode.SystemError, statuscode.GetText(statuscode.SystemError))
	}

	return response.NewSuccessRespMsg(token)
}

func (u *userService) Register(req request.RegisterReq) response.ResponseMsg {
	user := &model.User{
		UserBase: req.UserBase,
	}

	if err := u.userDao.Create(user); err != nil {
		global.LOG.Error("register user error", zap.Error(err))
		return response.NewErrorRespMsg(statuscode.RegisterAccountError, statuscode.GetText(statuscode.RegisterAccountError))
	}

	return response.NewSuccessRespMsg("")
}
