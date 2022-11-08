package middleware

import (
	"errors"
	"mychatroom/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//<----------用于gin的跨域鉴权中间件--------->

var JwtKey []byte

// 将设置中的jwtkey存入全局变量使用
func init() {
	jwtkey := viper.GetString("server.jwtkey")
	JwtKey = []byte(jwtkey)
}

type MyClaims struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

// 定义错误
var (
	ErrTokenExpired     error = errors.New("Token已过期,请重新登录")
	ErrTokenNotValidYet error = errors.New("Token验证失败,请重新登录")
	ErrTokenMalformed   error = errors.New("Token不正确,请重新登录")
	ErrTokenNotExist    error = errors.New("Token不存在,请重新登录")
	ErrTokenInvalid     error = errors.New("Token失效,请重新登录")
)

// CreateToken 生成token
func CreateToken(account, email string) (string, error) {
	claims := &MyClaims{
		Account: account,
		Email:   email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,    //签发时间
			ExpiresAt: time.Now().Unix() + 604800, //过期时间
			Issuer:    "LMY",                      //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// ParserToken 解析token
func ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	}

	return nil, ErrTokenNotExist
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_INEXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ERROR_TOKEN_INEXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 解析token
		claims, err := ParserToken(checkToken[1])
		if err != nil {
			if err == ErrTokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  errmsg.ERROR_TOKEN_RUNTIME,
					"message": errmsg.GetErrMsg(errmsg.ERROR_TOKEN_RUNTIME),
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Set("user_claims", claims)
		c.Next()
	}
}
