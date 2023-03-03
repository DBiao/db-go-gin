package dao

import (
	"db-go-gin/internal/app/model"
	"db-go-gin/internal/global"
)

type IUserDao interface {
	GetByUserName(username string) (*model.User, error)
	GetUserById(id uint32) (*model.User, error)
	GetAllUser(limit, offset int, data map[string]interface{}) ([]*model.User, error)
	GetUserCount(data map[string]interface{}) (int64, error)
	Create(user *model.User) error
	Update(data map[string]interface{}, id uint32) error
	Delete(id uint32) error
}

type userDao struct {
}

func NewUserDao() IUserDao {
	return &userDao{}
}

// GetByUserName 根据名称查找用户数据
func (u *userDao) GetByUserName(account string) (*model.User, error) {
	user := &model.User{}
	err := global.DB.Where("account = ?", account).First(&user).Error
	return user, err
}

// GetUserById 根据玩家Id查找用户数据
func (u *userDao) GetUserById(id uint32) (*model.User, error) {
	user := &model.User{}
	err := global.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// GetAllUser 获取所有用户
func (u *userDao) GetAllUser(limit, offset int, data map[string]interface{}) ([]*model.User, error) {
	db := global.DB.Limit(limit).Offset(offset)
	for key, value := range data {
		db.Where(key, value)
	}

	var users []*model.User
	err := db.Find(&users).Error
	return users, err
}

func (u *userDao) GetUserCount(data map[string]interface{}) (int64, error) {
	db := global.DB.Model(&model.User{})
	for key, value := range data {
		db.Where(key, value)
	}

	var total int64
	err := db.Count(&total).Error
	return total, err
}

// Create 创建用户
func (u *userDao) Create(user *model.User) error {
	return global.DB.Model(&model.User{}).Create(user).Error
}

// Update 修改用户
func (u *userDao) Update(data map[string]interface{}, id uint32) error {
	return global.DB.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
}

// Delete 删除用户
func (u *userDao) Delete(id uint32) error {
	user := model.User{}
	return global.DB.Model(&model.User{}).Where("id = ?", id).Delete(&user).Error
}
