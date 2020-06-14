package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shiniao/gtodo/pkg/errno"
	"github.com/shiniao/gtodo/pkg/log"
)

func Get()  {
	
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
	if err := c.ShouldBindJSON(&req); err != nil{
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
	if req.Email != req.ConfirmPassword{
		log.Warn("The twice password not match")
		SendResponse(c, errno.ErrTwicePasswordNotMatch, nil)
	}

	SendResponse(c, nil,nil)


}