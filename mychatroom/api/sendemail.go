package api

import (
	"context"
	"log"
	"mychatroom/dao"
	"mychatroom/service"
	"mychatroom/utils"
	"mychatroom/utils/errmsg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 发送邮箱验证码
func SendEmail(c *gin.Context) {
	email := c.PostForm("email")
	//检查email
	code := service.CheckEmail(email)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//发送邮件验证码
	verificationCode := utils.GetVerificationCode()
	err := utils.SendEmail(email, verificationCode)
	if err != nil {
		code = errmsg.ERROR_EMAIL_FAIL_TO_SEND
		log.Println("发送邮件失败", err)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//将随机生成的验证码存入redis
	expiretime := viper.GetInt64("server.expiretime")
	prefix := viper.GetString("server.redisprefix")
	if err = dao.Rdb.Set(context.Background(), prefix+email, verificationCode, time.Second*time.Duration(expiretime)).Err(); err != nil {
		code = errmsg.ERROR_REDIS
		log.Println("redis 存储验证码错误", err)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//返回操作成功信息
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
