package router

import (
	_ "db-go-gin/docs"
	"db-go-gin/internal/app/controller"
	"db-go-gin/internal/global"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var (
	userController controller.IUserController
)

func InitRouter() {
	// 初始化controller
	initController()
	// 初始化web
	newRouter()
}

// initController 初始化
func initController() {
	userController = controller.NewUserController()
}

// NewRouter 路由配置
func newRouter() *gin.Engine {
	// 创建gin
	router := newGin()

	// 使用跨域中间件和权限验证
	router.Use(cors.Default())
	//router.Use(middleware.Cors())

	// 开启pprof
	pprof.Register(router)

	// 验证码
	router.GET("math", controller.CaptchaMath)

	// swagger接口文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("ws", WsHandle)

	// 用户接口
	userGroup := router.Group("user")
	GroupApp.InitUserRouter(userGroup)

	// 启动http服务
	global.SERVER = initServer(fmt.Sprintf(":%d", global.CONFIG.System.Port), router)

	return nil
}

// 创建gin
func newGin() *gin.Engine {
	// gin调式模式
	mode := global.CONFIG.System.Mode
	if mode == gin.DebugMode {
		router := gin.Default()
		return router
	} else {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.Use(gin.Recovery()) // panic之后自动恢复
		router.Use(ginzap.Ginzap(global.LOG, time.RFC3339, true))
		router.Use(ginzap.RecoveryWithZap(global.LOG, true))
		return router
	}
}

// initServer 启动http服务
func initServer(address string, router *gin.Engine) *http.Server {
	s := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Error("http server start error", zap.Error(err))
			log.Fatal(err)
		}
	}()

	return s
}
