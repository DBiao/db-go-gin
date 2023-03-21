package cache

import (
	"github.com/go-redis/redis"
	"testing"
)

func Test(t *testing.T) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
	})

	_, err := client.Ping().Result()
	if err != nil {
		return
	}

	return
}
