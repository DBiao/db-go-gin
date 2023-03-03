package model

import "db-go-gin/internal/global"

type User struct {
	global.MODEL
	UserBase
}

type UserBase struct {
	Account      string `gorm:"type:varchar(64);uniqueIndex;comment:账号;not null" json:"account"` // 账号
	Password     string `gorm:"type:varchar(64);comment:密码;not null" json:"password"`            // 密码
	ContactName  string `gorm:"type:varchar(64);comment:联系人姓名;not null" json:"contact_name"`     // 联系人姓名
	MobileNumber string `gorm:"type:char(11);comment:手机号码;not null" json:"mobile_number"`        // 手机号码
	Mail         string `gorm:"type:varchar(100);comment:邮箱;not null" json:"mail"`               // 邮箱
	Unit         string `gorm:"type:varchar(100);comment:用户单位;not null" json:"unit"`             // 用户单位
	HeadPortrait string `gorm:"type:MediumBlob;comment:头像" json:"head_portrait"`                 //头像
	Name         string `gorm:"type:varchar(64);comment:名称;not null" json:"name"`                // 名称
}

func (u *User) TableName() string {
	return "user"
}
