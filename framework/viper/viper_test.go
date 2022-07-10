package viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

func TestViper(t *testing.T) {
	v := viper.New()
	// 路径必须要写相对路径,相对于项目的路径
	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 单个取值
	name := v.Get("name").(string)
	fmt.Println(v.Get("redis.host"))

	// 映射到结构体
	var s ServerConfig
	if err := v.Unmarshal(&s); err != nil {
		panic(err)
	}
	fmt.Printf("s: %+v\n", s)
	fmt.Println(name)
}
