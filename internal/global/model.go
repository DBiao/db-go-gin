package global

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        uint32         `gorm:"primarykey" json:"id"`    // 主键ID
	CreatedAt time.Time      `gorm:"index" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"index" json:"-"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`          // 删除时间
}
