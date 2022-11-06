package api

import (
	"log"
	"mychatroom/middleware"
	"mychatroom/model"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 添加用户到群聊
func AddUserToRoom(c *gin.Context) {
	account := c.PostForm("account") //想要添加的用户
	number := c.Param("number")      //想邀请进的指定聊天室

	if account == "" || number == "" {
		code := errmsg.ERROR_PARAM_EMPTY
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//检查邀请的账号是否存在
	b := service.IsAccountExist(account)
	if !b {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "你所发送的邀请号不属于你",
		})
		return
	}

	//检查邀请的聊天室号码是否属于用户自己
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	b = service.IsRoomBelongMine(mc.Account, number)
	if !b {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "你所发送的邀请号不属于你",
		})
		return
	}

	//检查是否为已经在本聊天室中
	if service.IsUserJoin(account, mc.Account) {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "已在此群聊中,无需添加",
		})
		return
	}
	// 保存自己与房间的关联记录
	my := &model.UserRoom{
		Account:   mc.Account,
		Number:    number,
		RoomType:  2,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := service.InsertOneUserRoom(my); err != nil {
		log.Printf("[mongoDB ERROR]:%v\n", err)
		code := errmsg.ERROR_MONGODB
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "数据库异常",
		})
		return
	}
	//保存被邀请者与房间的关联记录
	inviter := &model.UserRoom{
		Account:   account,
		Number:    number,
		RoomType:  2,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	if err := service.InsertOneUserRoom(inviter); err != nil {
		log.Printf("[mongoDB ERROR]:%v\n", err)
		code := errmsg.ERROR_MONGODB
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "数据库异常",
		})
		return
	}

	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code) + "添加好友成功",
	})

}

// 添加好友(添加到一个私人聊天室中，即互为好友)
func UserAdd(c *gin.Context) {
	account := c.PostForm("account") //想要添加的用户
	number := c.Param("number")      //想邀请进的指定聊天室

	if account == "" || number == "" {
		code := errmsg.ERROR_PARAM_EMPTY
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	//检查邀请的账号是否存在
	b := service.IsAccountExist(account)
	if !b {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "你所发送的邀请号不属于你",
		})
		return
	}

	//检查邀请的聊天室号码是否属于用户自己
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	b = service.IsRoomBelongMine(mc.Account, number)
	if !b {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "你所发送的邀请号不属于你",
		})
		return
	}

	//检查是否为好友
	if service.JudgeUserIsFriend(account, mc.Account) {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "互为好友,已在本聊天室中,不可重复添加",
		})
		return
	}

	// 保存自己与房间的关联记录
	my := &model.UserRoom{
		Account:   mc.Account,
		Number:    number,
		RoomType:  1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := service.InsertOneUserRoom(my); err != nil {
		log.Printf("[mongoDB ERROR]:%v\n", err)
		code := errmsg.ERROR_MONGODB
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "数据库异常",
		})
		return
	}
	//保存被邀请者与房间的关联记录
	inviter := &model.UserRoom{
		Account:   account,
		Number:    number,
		RoomType:  1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	if err := service.InsertOneUserRoom(inviter); err != nil {
		log.Printf("[mongoDB ERROR]:%v\n", err)
		code := errmsg.ERROR_MONGODB
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "数据库异常",
		})
		return
	}

	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code) + "添加好友成功",
	})
}

// 删除好友(从私人聊天室中移除)
func UserDelete(c *gin.Context) {
	account := c.Query("account")
	roomType, _ := strconv.Atoi(c.Query("room_type"))

	if account == "" {
		code := errmsg.ERROR_PARAM_EMPTY
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	// 获取房间账户关联
	number := service.GetUserRoomAccount(account, mc.Account, roomType)
	if number == "" {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "不为好友,无需删除",
		})
		return
	}

	// 删除user_room关联关系
	if err := service.DeleteUserRoom(number); err != nil {
		code := errmsg.ERROR_MONGODB
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "系统异常",
		})
		return
	}

	// 删除该私聊房间
	code := service.DeleteRoom(number)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 注销用户
func DeleteUser(c *gin.Context) {
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	//检查自己是否仍存在房间，若存在，需要删掉所有房间
	code := service.CheckOwnRoom(mc.Account)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "您仍有房间存在,需要删除全部房间才可以注销",
		})

		return
	}
	//删除用户
	code = service.DeleteUser(mc.Account)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 更改个人资料 (只能修改非固定信息)
func EditUserInfo(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)
	code := service.EditUserInfo(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改密码
func ChangePassword(c *gin.Context) {
	password := c.PostForm("password")
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	code := service.ChangePassword(mc.Account, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户个人信息
func UserInfo(c *gin.Context) {
	v, _ := c.Get("user_claims")
	mc := v.(*middleware.MyClaims)
	data, code := service.GetUserInfoByAccount(mc.Account)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询指定用户信息
func UserQuery(c *gin.Context) {

	account := c.Query("account")
	if account == "" {
		code := errmsg.ERROR_PARAM_EMPTY
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	user, code := service.GetUserInfoByAccount(account)
	if code != errmsg.SUCCESS {
		log.Println("查询信息失败")
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + "查询数据异常",
		})
		return
	}

	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	data := model.UserQuery{
		Account:  user.Account,
		Nickname: user.Nickname,
		Sex:      user.Sex,
		Avatar:   user.Avatar,
		IsFriend: false,
	}
	//判断是否为好友(在同一私聊聊天室内)
	if service.JudgeUserIsFriend(user.Account, mc.Account) {
		data.IsFriend = true
	}

	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}
