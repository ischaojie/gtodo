/*user_repo 直接对接数据库，提供基础的CURD功能开放给handler*/
package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/shiniao/gtodo/internal/cache"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/pkg/log"
)

// UserRepo 定义用户仓库接口
type Repo interface {
	Create(db *gorm.DB, user model.UserModel) (id uint64, err error)
	Update(db *gorm.DB, id uint64, userMap map[string]interface{}) error
	GetUserByID(db *gorm.DB, id uint64) (*model.UserModel, error)
	GetUserByEmail(db *gorm.DB, email string) (*model.UserModel, error)
}

type UserRepo struct {
	userCache *cache.Cache
}

// Create 向数据库插入新用户
func (ur *UserRepo) Create(db *gorm.DB, user model.UserModel) (id uint64, err error) {
	err = db.Create(&user).Error
	if err != nil {
		return 0, errors.Wrap(err, "[user repo] create user err")
	}
	return user.ID, nil
}

// Update 更新用户信息
func (ur *UserRepo) Update(db *gorm.DB, id uint64, userMap map[string]interface{}) error {
	user, err := ur.GetUserByID(db, id)
	if err != nil {
		return errors.Wrap(err, "[user_repo] update user data err")
	}
	// 删除user cache
	if err = ur.userCache.DelUserCache(id); err != nil {
		log.Warnf("[user_repo] delete user cache err: %v", err)
	}
	return db.Model(user).Updates(userMap).Error
}

// GetUserByID 通过ID获取用户信息
func (ur *UserRepo) GetUserByID(db *gorm.DB, id uint64) (*model.UserModel, error) {
	panic("implement me")
}

func (ur *UserRepo) GetUserByEmail(db *gorm.DB, email string) (*model.UserModel, error) {
	panic("implement me")
}
