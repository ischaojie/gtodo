package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shiniao/gtodo/handler"
	"github.com/shiniao/gtodo/pkg/errno"
	"github.com/shiniao/gtodo/pkg/token"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

