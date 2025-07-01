package core

import (
	"gin-admin-api/config"
	"gin-admin-api/global"

	"github.com/go-redis/redis/v8"
)

// redis配置初始化
 
var (
	RedisDb *redis.Client
)


func RedisInit()error{
	RedisDb = redis.NewClient(&redis.Options{
		Addr:config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB: config.Config.Redis.Db,
	})
	_,err := RedisDb.Ping(global.Ctx).Result()
	if err != nil{
		return err
	}
	global.Log.Infof("[redis]連接成功")
	return nil
}