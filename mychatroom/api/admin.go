package api

import (
	"mychatroom/model"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 管理员登录
func AdminLogin(c *gin.Context) {
	var admin model.Admin
	_ = c.ShouldBind(&admin)
	code := service.CheckAdmin(admin.Username, admin.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    admin.Username,
		"message": errmsg.GetErrMsg(code),
	})
}

// 管理员删除用户
func AdminDeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	//根据要删除的用户id检索账号
	account, code := service.GetUserAccountById(id)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	//利用检索的账号查询该账号名下是否有房间存在
	code = service.CheckOwnRoom(account)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "该账号名下仍有房间存在,需要删除全部房间才可以注销",
		})

		return
	}

	//通过检测，删除账号
	code = service.AdminDeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 管理员删除聊天室
func AdminDeleteRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.AdminDeleteRoom(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 管理员获得用户列表
func GetUserList(c *gin.Context) {
	data, code := service.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
