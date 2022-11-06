package api

import (
	"context"
	"log"
	"mychatroom/dao"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 忘记密码功能，可以通过发送邮箱验证重置密码
func ForgetPassword(c *gin.Context) {
	account := c.PostForm("account")
	email := c.PostForm("email")
	password := c.PostForm("password")
	vfcode := c.PostForm("vfcode")
	//检查账号是否存在
	code := service.CheckAccount(account)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 检查验证码是否正确
	prefix := viper.GetString("server.redisprefix")
	r, err := dao.Rdb.Get(context.Background(), prefix+email).Result()
	if err != nil {
		code = errmsg.ERROR_REDIS
		log.Println("redis获取验证码错误", err)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "获取验证码失败",
		})
		return
	} else if r != vfcode {
		code = errmsg.ERROR_CAPTCHA_WRONG
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	code = service.ChangePassword(account, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
