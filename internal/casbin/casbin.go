package source

import (
	"db-go-gin/internal/global"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

const ()

var carbines = []gormadapter.CasbinRule{}

// InitCasbin 初始化权限表数据
func InitCasbin() error {
	if err := global.DB.AutoMigrate(gormadapter.CasbinRule{}); err != nil {
		return err
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		// casbin_rule表中有权限数据，则不再插入权限数据，加入新的权限，需要把casbin_rule表数据全部清空
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected > 0 {
			color.Info.Println("[Mysql] --> casbin_rule 权限表的权限数据已存在!")
			return nil
		}

		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}

		color.Info.Println("[Mysql] --> casbin_rule 权限表初始化数据成功!")
		return nil
	})
}
