/*user_repo 直接对接数据库，提供基础的CURD功能开放给handler*/
package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/shiniao/gtodo/internal/cache"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/pkg/log"
	"github.com/shiniao/gtodo/pkg/redis"
	"time"
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

func NewUSerRepo() Repo {
	return &UserRepo{userCache: cache.NewUserCache()}
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
	// 首先看cache里面有没有
	userModel, err := ur.userCache.GetUserCache(id)
	if err != nil {
		return nil, errors.Wrapf(err, "[user_repo] get user cache data err")
	}
	// 有cache，返回cache的userModel
	if userModel != nil && userModel.ID != 0 {
		return userModel, nil
	}

	// 从数据库获取
	// 防止缓存击穿，加锁
	key := fmt.Sprintf("uid:%d", id)
	lock := redis.NewLock(redis.Client, key, 3*time.Second)
	token := lock.GenToken()

	isLock, err := lock.Lock(token)
	if err != nil || !isLock {
		return nil, errors.Wrapf(err, "[user_repo] lock err")
	}
	defer lock.UnLock(token)

	data := &model.UserModel{}
	if isLock {
		// 从数据库查找
		err = db.Where(&model.UserModel{ID: id}).First(data).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, "[user_repo] get user data err")
		}
		// 写入cache
		err = ur.userCache.SetUserCache(id, data)
		if err != nil {
			return data, errors.Wrap(err, "[user_repo] set user data err")
		}
	}
	return data, nil
}

func (ur *UserRepo) GetUserByEmail(db *gorm.DB, email string) (*model.UserModel, error) {
	panic("implement me")
}
