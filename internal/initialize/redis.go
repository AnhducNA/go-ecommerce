package initialize

import (
	"context"
	"fmt"

	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // No password set
		DB:       r.Database, // Use default DB
		Protocol: 2,          // Connection protocol
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("InitRedis error", zap.Error(err))
	}

	fmt.Println("InitRedis success")
	global.Rdb = rdb
	RedisExample()
}

func RedisExample() {
	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := global.Rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key", val)
}
