package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 用于外部操作
var Mdb *mongo.Database

func init() {
	//获取配置
	driver := viper.GetString("mdb.driver")
	host := viper.GetString("mdb.host")
	port := viper.GetString("mdb.port")
	dbname := viper.GetString("mdb.dbname")

	dsn := fmt.Sprintf("%s://%s:%s", driver, host, port)

	//超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//连接mongodb客户端
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal("连接 MongoDB 失败:", err)
	}

	//打开数据库
	mdb := client.Database(dbname)

	//连接外部暴露的出口
	Mdb = mdb
}
