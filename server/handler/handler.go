package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"mini_todo/errno"
	"mini_todo/token"
	"net/http"
)

/*统一api返回内容*/

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// * 统一返回给client的内容
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}

type Key struct {
	Key string `json:"key"`
}

func Token(c *gin.Context) {
	var key Key

	c.BindJSON(&key)
	log.Print(key)
	//key := c.PostForm("key")
	// * 判断key是否正确
	if key.Key != viper.GetString("key") {
		SendResponse(c, errno.ErrKeyIncorrect, nil)
		return
	}
	// * Sign the json web token.
	t, err := token.Sign(c,key.Key, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, token.Token{Token: t})
}

