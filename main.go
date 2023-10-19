package main

import (
	"gin-z/bootstrap"
	"gin-z/global"
)

// @title         项目入口，一切的开始
// @description   进行项目所需要配置初始化，并启动服务器
func main() {
	// 初始化全局配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 启动服务器
	bootstrap.RunServer()
}
