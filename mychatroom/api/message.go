package api

import (
	"mychatroom/middleware"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查看聊天记录
func MessageList(c *gin.Context) {
	var code int
	number := c.Query("number")
	pageNum, _ := strconv.ParseInt(c.Query("page_num"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 32)
	skip := (pageNum - 1) * pageSize

	if number == "" {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"stauts":  code,
			"message": errmsg.GetErrMsg(code) + "房间号不能为空",
		})
		return
	}

	//判断房间是否存在
	code = service.CheckRoom(number)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"stauts":  code,
			"message": errmsg.GetErrMsg(code) + "房间不存在,请重新输入房间号",
		})
		return
	}

	//判断用户是否属于该房间
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	_, err := service.GetUserRoomByAccountNumber(mc.Account, number)
	if err != nil {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"stauts":  code,
			"message": errmsg.GetErrMsg(code) + "非法访问",
		})
		return
	}

	//聊天记录查询
	data, err := service.GetMessageListByRid(number, &pageSize, &skip)
	if err != nil {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + err.Error(),
		})
		return
	}
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    data,
	})
}
