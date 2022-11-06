package api

import (
	"context"
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 注册账号
func Register(c *gin.Context) {
	var data model.User
	var code int
	vfcode := c.Query("vfcode")
	_ = c.ShouldBind(&data)

	if vfcode == "" || data.Email == "" || data.Account == "" || data.Password == "" {
		code = errmsg.ERROR_PARAM_EMPTY
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	// 验证码是否正确
	prefix := viper.GetString("server.redisprefix")
	r, err := dao.Rdb.Get(context.Background(), prefix+data.Email).Result()
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

	// 判断账号是否唯一
	code = service.CheckAccount(data.Account)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//注册一个账号
	code = service.Register(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"account":  data.Account,
			"email":    data.Email,
			"Nickname": data.Nickname,
		},
		"message": errmsg.GetErrMsg(code),
	})
}
