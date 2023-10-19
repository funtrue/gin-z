// @Title  项目启动初始化
// @Description  基本信息初始化,获取config.yaml文件内容并进行初始化
package bootstrap

import (
	"fmt"
	"gin-z/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// @title    全局配置文件初始化函数
// @description   初始化config.yaml里面的所有内容，赋值给全局变量
// @param     输入参数名:NULL              参数类型:NULL
// @return    返回参数名:vip_config        参数类型:Viper
func InitializeConfig() *viper.Viper {
	// 设置配置文件路径
	config_address := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if ConfigEnv := os.Getenv("VIPER_CONFIG_ADDRESS"); ConfigEnv != "" {
		config_address = ConfigEnv
	}

	// 初始化 viper
	vip_config := viper.New()
	vip_config.SetConfigFile(config_address)
	vip_config.SetConfigType("yaml")
	if err := vip_config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件
	vip_config.WatchConfig()
	vip_config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := vip_config.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量
	if err := vip_config.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return vip_config
}
