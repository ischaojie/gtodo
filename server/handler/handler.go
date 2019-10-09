package handler

import (
	"github.com/gin-gonic/gin"
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

func Token(c *gin.Context) {
	// Sign the json web token.
	t, err := token.Sign(c,"")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, token.Token{Token: t})
}

