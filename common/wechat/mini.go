package wechat

import (
	"gin-example/config"
	"gin-example/pkg/logger"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
)

func initMiniApp() {
	wxConfig := config.WxConfig{}

	_, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  wxConfig.AppId,
		Secret: wxConfig.Secret, // 小程序app secret
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       0,
		}),
		//HttpDebug: true,
		Log: miniProgram.Log{
			Level: "info",
			File:  "./logs/wechat.log",
		},
	})
	if err != nil {
		logger.Logger.Panic("初始化微信错误")
	}
	logger.Logger.Debug("微信小程序初始化完成")
}
