package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/internal/repository"
	"github.com/shiniao/gtodo/pkg/auth"
	token2 "github.com/shiniao/gtodo/pkg/token"
	"time"
)

// Service 代表服务层
type Service interface {
	Register(ctx *gin.Context, username, email, password string) error             // 注册新用户
	EmailLogin(ctx *gin.Context, email, password string) (token string, err error) // 邮箱登录
	PhoneLogin(ctx *gin.Context, phone, code int) (token string, err error)        // 手机登录
	GetUserByID(id uint64) (*model.UserModel, error)                               // 根据用户ID获取用户
	GetUserInfoByID(id uint64) (*model.UserInfo, error)                            // 根据用户ID获取用户信息
	GetUserByEmail(email string) (*model.UserModel, error)                         // 根据邮箱获取用户
}

type userService struct {
	userRepo repository.Repo
}

// Register 注册新用户
func (s *userService) Register(ctx *gin.Context, username, email, password string) error {
	pwd, err := auth.Encrypt(password)
	if err != nil {
		return errors.Wrapf(err, "Encrypt password err.")
	}
	u := model.UserModel{
		Username:  username,
		Password:  pwd,
		Email:     email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err = s.userRepo.Create(model.GetDB(), u)
	if err != nil {
		return errors.Wrapf(err, "Create user err")
	}
	return nil

}

// EmailLogin 用邮箱登录
func (s *userService) EmailLogin(c *gin.Context, email, password string) (tokenString string, err error) {
	u, err := s.GetUserByEmail(email)
	if err != nil {
		return "", errors.Wrapf(err, "get user by email err")
	}
	// 密码是否正确
	err = auth.Compare(u.Password, password)
	if err != nil {
		return "", errors.Wrapf(err, "password compare err")
	}
	// 签发token
	tokenString, err = token2.Sign(c, token2.Context{UserID: u.ID, UserName: u.Username}, "")
	if err != nil {
		return "", errors.Wrapf(err, "sign token err")
	}

	return tokenString, nil

}

// PhoneLogin 用手机登陆
func (s *userService) PhoneLogin(ctx *gin.Context, phone, code int) (token string, err error) {
	return "", nil
}

// GetUserByID 通过id获取用户
func (s *userService) GetUserByID(id uint64) (*model.UserModel, error) {
	userModel, err := s.userRepo.GetUserByID(model.GetDB(), id)
	if err != nil {
		return userModel, errors.Wrapf(err, "get user info err from db by id: %d", id)
	}

	return userModel, nil
}

// GetUserInfoByID 通过id获取用户信息
func (s *userService) GetUserInfoByID(id uint64) (*model.UserInfo, error) {
	// user, err := s.userRepo.GetUserByID(model.GetDB(), id)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "[user_service] get user err")
	// }
	// // UserModel 转 UserInfo
	// return user, nil
	return nil, nil
}

func (userService) GetUserByEmail(email string) (*model.UserModel, error) {
	panic("implement me")
}

// NewUserService 初始化user service
func NewUserService() Service {
	return &userService{userRepo: repository.NewUSerRepo()}
}

// UserSvc 对外提供Service接口
var UserSvc = NewUserService()
