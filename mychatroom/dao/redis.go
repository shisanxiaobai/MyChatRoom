package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 外部调用接口
var Rdb *redis.Client

// redis初始化
func init() {
	//获取配置
	addr := viper.GetString("rdb.addr")
	loc := viper.GetInt("rdb.local")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",  //redis密码，没密码可设为空或者不写这个属性
		DB:       loc, //数据库位置，放在几号数据库，默认是0号
	})

	//连接暴露的出口
	Rdb = rdb
}
