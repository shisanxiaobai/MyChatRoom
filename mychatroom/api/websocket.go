package api

import (
	"log"
	"mychatroom/middleware"
	"mychatroom/model"
	"mychatroom/service"
	"mychatroom/utils/errmsg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var AccountConn = make(map[string]*websocket.Conn) //根据用户名存储websocket连接

func WebsocketMessage(c *gin.Context) {
	//将http歇息升级为websocket协议，返回一个websocket连接
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		log.Println("websocket协议升级错误", err)
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code) + err.Error(),
		})
		return
	}

	defer conn.Close()
	//获取jwt中设置的token令牌，转成MyClamis类型
	mc := c.MustGet("user_claims").(*middleware.MyClaims)
	//map存储账号和账号的连接
	AccountConn[mc.Account] = conn

	//循环读取消息
	for {
		var im *model.InstantMessage //一个消息实例
		err := conn.ReadJSON(&im)
		if err != nil {
			log.Println("读取消息错误", err)
			return
		}

		//判断用户是否属于消息体的房间
		_, err = service.GetUserRoomByAccountNumber(mc.Account, im.Number)
		if err != nil {
			log.Println(mc.Account, "+", im.Number, "此消息所属房间不存在", err)
			return
		}

		//保存消息
		message := &model.Message{
			Account:   mc.Account,
			Number:    im.Number,
			Data:      im.Message,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		//将消息保存到mongodb
		err = service.InsertMessage(message)
		if err != nil {
			log.Println("mongodb插入数据失败", err)
			return
		}

		//获取指定房间的在线用户
		userRooms, err := service.GetUserRoomByNumber(message.Number)
		if err != nil {
			log.Println("查询room出错", err)
			return
		}
		for _, userRoom := range userRooms {
			if wsConn, ok := AccountConn[userRoom.Account]; ok {
				err = wsConn.WriteMessage(websocket.TextMessage, []byte(im.Message))
				if err != nil {
					log.Println("WriteMessage 错误", err)
					return
				}
			}
		}
	}
}
