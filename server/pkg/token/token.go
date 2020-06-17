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

type Context struct {
	UserID   uint64
	UserName string
}

var ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")

// Sign 下发token
func Sign(c *gin.Context, c Context, secret string) (tokenString string, err error) {

	// 从 config 读取secret
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// jwt claims
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		// 过期时间
		// "exp":      time.Now().Add(time.Hour * 2).Unix(),
		"uid":   c.userID,
		"uname": c.userName,
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}

// Parse 判断token合法性，并返回token中的用户信息
func Parse(token, secret string) (*Context, error) {
	ctx := &Context{}

	// parse the token
	t, err := jwt.Parse(token, secretFunc(secret))
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		ctx.UserID = uint64(claims["uid"].(float64))
		ctx.UserName = claims["uname"].(string)
		return ctx, nil

	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// * Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// * Parse the header to get the token part.
	_, err := fmt.Sscanf(header, "Bearer %s", &t)
	if err != nil {
		fmt.Printf("fmt.Sscanf err,: %+v", err)
	}
	return Parse(t, secret)
}

// secretFunc 验证密钥的格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// * Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
