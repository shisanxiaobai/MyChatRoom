package api

import (
	"mychatroom/middleware"
	"mychatroom/model"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建一个聊天室
func CteateRoom(c *gin.Context) {
	number := c.PostForm("number")
	name := c.PostForm("name")
	info := c.PostForm("info")
	mc := c.MustGet("user_claims").(*middleware.MyClaims)

	if number == "" {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "聊天室号码不能为空",
		})
		return
	}

	code := service.CheckRoom(number)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	code = service.CreateRoom(mc.Account, number, name, info)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除一个聊天室
func DeleteRoom(c *gin.Context) {
	number := c.Param("number")
	code := service.DeleteRoom(number)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改聊天室的信息(只能修改非固定信息)
func EditRoom(c *gin.Context) {
	var data model.Room
	number := c.Param("number")
	_ = c.ShouldBind(&data)
	code := service.EditRoom(number, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查看自己的聊天室信息
func OwnRoom(c *gin.Context) {
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	data, code := service.OwnRoom(mc.Account)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查看所有房间列表
func RoomList(c *gin.Context) {
	data, code := service.RoomList()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
