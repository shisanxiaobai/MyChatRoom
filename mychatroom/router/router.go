package router

import (
	"mychatroom/api"
	"mychatroom/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 路由入口
func RouterStart() {
	//配置获取
	port := viper.GetString("server.httpport")

	r := gin.New()

	//中间件的使用
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())

	router := r.Group("/api")
	{
		//用户登录
		router.POST("/login", api.Login)
		//发送邮箱验证码
		router.POST("/send/email", api.SendEmail)
		//用户注册
		router.POST("/register", api.Register)
		//忘记密码
		router.POST("/forget", api.ForgetPassword)
	}

	// <-----------------需要token-------------->

	auth := r.Group("/api/u")
	{
		auth.Use(middleware.JwtToken()) //auth使用jwt中间件

		//用户模块操作
		auth.POST("/user/add/group/:number")                       //创建群聊
		auth.POST("/user/add/friend/:number", api.UserAdd)         //添加好友
		auth.DELETE("/user/delete/friend/:number", api.UserDelete) //删除好友
		auth.DELETE("/user/delete", api.DeleteUser)                //注销用户
		auth.PUT("/user/edit/:id", api.EditUserInfo)               //修改个人资料
		auth.PUT("user/password", api.ChangePassword)              //修改密码
		auth.GET("/user/info", api.UserInfo)                       //查看个人资料
		auth.GET("/user/query", api.UserQuery)                     //查询指定用户

		//消息模块操作
		auth.GET("/websocket/message", api.WebsocketMessage) //发送，接受消息
		auth.GET("/message/list", api.MessageList)           //聊天记录列表

		//聊天室模块操作
		auth.POST("/room/add", api.CteateRoom)
		auth.DELETE("/room/delete/:number", api.DeleteRoom)
		auth.PUT("/room/edit/:number", api.EditRoom)
		auth.GET("room/Own", api.OwnRoom)
		auth.GET("room/list", api.RoomList)
	}

	//<---------------系统管理员----------------------->

	admin := r.Group("/api/admin")
	{
		admin.POST("/login", api.AdminLogin)
		admin.DELETE("/userdelete/:id", api.AdminDeleteUser)
		admin.DELETE("/roomdelete/:id", api.AdminDeleteRoom)
		admin.GET("/userlist", api.GetUserList)
	}
	r.Run(port)
}
