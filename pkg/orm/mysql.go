package orm

import (
	"db-go-gin/internal/app/model"
	"db-go-gin/internal/global"
	"errors"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMysql 初始化Mysql数据库
func InitMysql() (*gorm.DB, error) {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil, errors.New("mysql name is empty")
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig())
	if err != nil {
		return db, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)

	return db, nil
}

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		return err
	}

	global.LOG.Info("mysql register table success")

	return nil
}

// gormConfig 根据配置决定是否开启日志
func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(logger.Info)
	default:
		config.Logger = Default.LogMode(logger.Info)
	}

	return config
}
