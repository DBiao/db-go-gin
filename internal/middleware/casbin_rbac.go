package middleware

import (
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/global"
	"db-go-gin/internal/global/statuscode"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*multi.MultiClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId

		// 判断策略中是否存在
		casbin, err := Casbin()
		if err != nil {
			global.LOG.Error("casbin is error", zap.Error(err))
			response.Response(c, response.NewErrorRespMsg(statuscode.SystemPermissionError, statuscode.GetText(statuscode.SystemPermissionError)))
			c.Abort()
			return
		}

		// 是否拥有权限
		success, err := casbin.Enforce(sub, obj, act)
		if err != nil {
			global.LOG.Error("enforce is error", zap.Error(err))
			response.Response(c, response.NewErrorRespMsg(statuscode.SystemPermissionError, statuscode.GetText(statuscode.SystemPermissionError)))
			c.Abort()
			return
		}
		if !success {
			response.Response(c, response.NewErrorRespMsg(statuscode.SystemPermissionError, statuscode.GetText(statuscode.SystemPermissionError)))
			c.Abort()
			return
		}
		c.Next()
	}
}

// Casbin 持久化到数据库  引入自定义规则
func Casbin() (*casbin.SyncedEnforcer, error) {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.CONFIG.System.ModelPath, a)

		syncedEnforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	})

	err := syncedEnforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return syncedEnforcer, nil
}

// ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
