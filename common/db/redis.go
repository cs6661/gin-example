package db

import (
	"context"
	"gin-example/pkg/logger"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func initRedis(ctx context.Context) {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	RDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	logger.Logger.Debug("redis初始化完成")
	ping := RDB.Ping(ctx)
	if ping.Err() != nil {
		logger.Logger.Sugar().Errorf("redis ping error: %v", ping.Err())
	}
}
