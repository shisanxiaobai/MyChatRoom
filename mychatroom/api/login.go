package api

import (
	"mychatroom/middleware"
	"mychatroom/model"
	"mychatroom/utils/errmsg"
	"net/http"

	"mychatroom/service"

	"github.com/gin-gonic/gin"
)

// @Summary	用户登录
// @Description 用户登录
// @Tags	Login
// @Accept  json
// @Produce json
// @Param 	login body model.User true "登录"
// @Success 200 {object} model.Reponse
// @Failure 500 {object} model.Reponse
// @Router /api/login [post]
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	code := service.CheckLogin(user.Account, user.Password)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	data, _ := service.GetUserInfoByAccount(user.Account)
	SetToken(c, user.Account, data.Email)
}

// 设置Token
func SetToken(c *gin.Context, account, email string) {

	token, err := middleware.CreateToken(account, email)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"data": gin.H{
			"account": account,
			"email":   email,
		},
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
		"token":   token,
	})
}
