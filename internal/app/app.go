package app

import (
	"context"
	cs "db-go-gin/internal/casbin"
	"db-go-gin/internal/global"
	"db-go-gin/internal/router"
	"db-go-gin/pkg/cache"
	"db-go-gin/pkg/logger"
	"db-go-gin/pkg/orm"
	"db-go-gin/pkg/shutdown"
	"db-go-gin/pkg/viper"
	"log"
	"time"

	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// Start 初始化服务
func Start() {
	var err error

	// 初始化配置
	if global.VIPER, err = viper.InitViper(); err != nil {
		log.Fatal(err)
	}

	// 初始化日志
	global.LOG = logger.InitZap()

	// 初始化jwt
	if err = multi.InitDriver(&multi.Config{
		DriverType:    "local",
		TokenMaxCount: 100,
	}); err != nil {
		log.Fatal(err)
	}

	// 初始化mysql
	if global.DB, err = orm.InitMysql(); err != nil {
		log.Fatal(err)
	}

	// 初始化mysql表
	if err = orm.MysqlTables(global.DB); err != nil {
		log.Fatal(err)
	}

	// 初始化权限表数据
	if err = cs.InitCasbin(); err != nil {
		log.Fatal(err)
	}

	// 初始化Gin
	router.InitRouter()
}

// Close 优雅关闭
func Close() {
	shutdown.NewHook().Close(
		// 关闭http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := global.SERVER.Shutdown(ctx); err != nil {
				global.LOG.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭mysql
		func() {
			if global.DB != nil {
				if db, err := global.DB.DB(); err != nil {
					if err = db.Close(); err != nil {
						global.LOG.Error("mysql close err", zap.Error(err))
					}
				}
			}
		},

		// 关闭Redis
		func() {
			cache.Close()
		})
}
