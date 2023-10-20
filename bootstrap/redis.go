package bootstrap

import (
	"context"
	"gin-z/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// @title    redis初始化函数
// @description  redis初始化全局配置
// @param     输入参数名:NULL              参数类型:NULL
// @return    返回参数名:client               参数类型:redis
func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
