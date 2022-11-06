package main

import (
	_ "mychatroom/conf"
	"mychatroom/router"
)

//<------------------项目启动入口----------------------->
//<----------下面文档注释用于生成swagger文档------------->

// @title mychatroom gin+gorm+redis+mongodb
// @version 1.0
// @description gin+gorm+redis+mongodb+jwt+cors+viper等等
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:9001
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Basepath /
func main() {
	router.RouterStart()
}
