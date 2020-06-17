package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shiniao/gtodo/internal/service"
	"github.com/shiniao/gtodo/pkg/errno"
	"github.com/shiniao/gtodo/pkg/log"
	"github.com/shiniao/gtodo/pkg/token"
)

func Get() {

}

// RegisterRequest 代表注册结构体
type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	// 判断注册字段
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("register bind param err: %v", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("register req: %#v", req)
	// 判断是否有对应字段
	if req.Username == "" || req.Password == "" || req.Email == "" {
		log.Warnf("params is empty: %v", req)
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	// 判断两次密码是否一致
	if req.Email != req.ConfirmPassword {
		log.Warn("The twice password not match")
		SendResponse(c, errno.ErrTwicePasswordNotMatch, nil)
	}

	err := service.UserSvc.Register(c, req.Username, req.Email, req.Password)
	if err != nil {
		log.Warnf("register err: %+v", err)
		SendResponse(c, errno.ErrRegister, nil)
	}

	SendResponse(c, nil, nil)

}

// LoginCredentials 默认登录方式-邮箱
type LoginCredentials struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginCredentials
	if err := c.Bind(&req); err != nil {
		log.Warnf("email login bind param err: %+v", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("login req %#v", req)
	// 判断是否有对应字段
	if req.Password == "" || req.Email == "" {
		log.Warnf("params is empty: %v", req)
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	// 默认邮箱登录
	tokenString, err := service.UserSvc.EmailLogin(c, req.Email, req.Password)
	if err != nil {
		log.Warnf("email login err: %v", err)
		SendResponse(c, errno.ErrEmailOrPassword, nil)
	}

	SendResponse(c, nil, token.Token{Token: tokenString})
}

// PhoneLoginCredentials 手机登录
type PhoneLoginCredentials struct {
	Phone int `json:"phone" form:"phone" binding:"required" example:"13010002000"`
	VCode int `json:"vcode" form:"vcode" binding:"required" example:"120110"`
}

// PhoneLogin 邮箱登录
func PhoneLogin(c *gin.Context) {
	var req PhoneLoginCredentials
	if err := c.Bind(&req); err != nil {
		log.Warnf("phone login bind param err: %+v", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("login req %#v", req)
	// 判断是否有对应字段
	if req.VCode == 0 || req.Phone == 0 {
		log.Warnf("params is empty: %v", req)
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	// 验证码检查
	if !service.VCodeSvc.CheckLoginVCode(req.Phone, req.VCode) {
		SendResponse(c, errno.ErrVCode, nil)
		return
	}

	// 登录
	tokenString, err := service.UserSvc.PhoneLogin(c, req.Phone, req.VCode)
	if err != nil {
		SendResponse(c, errno.ErrVCode, nil)
		return
	}
	SendResponse(c, nil, token.Token{Token: tokenString})
}

// GetUserID 返回用户id
func GetUserID(c *gin.Context) uint64 {
	if c == nil {
		return 0
	}

	// uid 必须和 middleware/auth 中的 uid 命名一致
	if v, exists := c.Get("uid"); exists {
		uid, ok := v.(uint64)
		if !ok {
			return 0
		}

		return uid
	}
	return 0
}

// Account 获取账户相关信息
func Account(c *gin.Context) {
	uid := GetUserID(c)
	user, err := service.UserSvc.GetUserByID(uid)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

func AccountUpdate(c *gin.Context) {

}
