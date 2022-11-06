package dao

import (
	"fmt"
	"log"
	"mychatroom/model"

	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 用于外部操作
var Db *gorm.DB

// mysql 初始化操作
func init() {
	//获取配置
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	charset := viper.GetString("db.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbname, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		//创建表时的命名策略
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("连接 mysql 失败", err)
	}

	//迁移数据表，若没有此表则自动创建
	db.AutoMigrate(&model.User{}, &model.Room{})

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//连接暴露的接口
	Db = db
}
