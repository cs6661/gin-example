package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string     `yaml:"name"`
	Mode      string     `yaml:"mode"`
	Version   string     `yaml:"version"`
	Mysql     *Mysql     `yaml:"mysql"`
	LogConfig *LogConfig `yaml:"log"`
}

type Mysql struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// LogConfig 日志信息
type LogConfig struct {
	Path        string     `yaml:"path"`
	LogFileName string     `yaml:"logFileName"`
	Stdout      bool       `yaml:"stdout"`
	Level       string     `yaml:"level"`
	LumberJack  LumberJack `yaml:"lumberJack"`
}

// LumberJack 日志切割
type LumberJack struct {
	MaxSize    int  `yaml:"maxSize"`    //单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/config/")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	// 把读取到的配置信息反序列化到Conf变量中
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}

	if err = viper.UnmarshalKey("log", &Conf.LogConfig); err != nil {
		panic(fmt.Errorf("unmarshal to LogConfig failed, err:%v", err))
	}

}
