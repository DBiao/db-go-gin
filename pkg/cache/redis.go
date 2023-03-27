package cache

import (
	"db-go-gin/internal/global"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

// InitRedis 初始化连接
func InitRedis() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         global.CONFIG.Redis.Host + ":" + global.CONFIG.Redis.Port,
		Password:     global.CONFIG.Redis.Password, // no password set
		DB:           global.CONFIG.Redis.DB,       // use default DB
		PoolSize:     global.CONFIG.Redis.PoolSize,
		MinIdleConns: global.CONFIG.Redis.MinIdleConn,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	if client != nil {
		_ = client.Close()
	}
}
