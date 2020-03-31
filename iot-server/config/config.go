package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var Conf *viper.Viper

var once sync.Once // 单例工具。

// 获取全局变量 ,单例模式。
func GetConf() *viper.Viper {
	once.Do(func() {
		if Conf == nil {
			Conf = ReadConfigFile()
		}
	})
	return Conf
}

//读取配置文件。
func ReadConfigFile() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")  // name of config file (without extension)
	v.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("../config/app/")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Print(v.AllKeys())
	return v
}