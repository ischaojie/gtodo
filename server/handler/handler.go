package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shiniao/gtodo/pkg/errno"
	"github.com/shiniao/gtodo/pkg/token"
	"github.com/spf13/viper"
	"net/http"
)

// Response api 返回结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse 统一返回给client的内容
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

type Key struct {
	Key string `json:"key"`
}

// RouteNotFound 返回路由不存在
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "the route not found")
}

func Token(c *gin.Context) {
	var key Key

	c.BindJSON(&key)
	// key := c.PostForm("key")
	// * 判断key是否正确
	if key.Key != viper.GetString("key") {
		SendResponse(c, errno.ErrKeyIncorrect, nil)
		return
	}
	// * Sign the json web token.
	t, err := token.Sign(c, key.Key, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, token.Token{Token: t})
}

