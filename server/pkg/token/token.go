package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

var ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")


// * 下发token
func Sign(ctx *gin.Context, key, secret string) (tokenString string, err error) {

	// * 读取config
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// * jwt claims
	// TODO 过期时间处理
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// * 用户申请token携带的key
		"key": key,
		// * 生效时间
		"nbf":      time.Now().Unix(),
		// * 签发时间
		"iat":      time.Now().Unix(),
		// * 过期时间
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	// * Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}


func ParseRequest(c *gin.Context) error {
	header := c.Request.Header.Get("Authorization")

	// * Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return ErrMissingHeader
	}

	var t string
	// * Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}


// * 解析token
func Parse(token, secret string) error {

	// * parse the token
	t, err := jwt.Parse(token, secretFunc(secret))

	// Parse error.
	if err != nil {
		return err

		// * Read the token if it's valid.
	} else if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return nil

	} else {
		return err
	}
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// * Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
