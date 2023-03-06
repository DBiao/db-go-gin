package middleware

import (
	"time"

	"db-go-gin/internal/global"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// keyFunc 在缓存中的key值
func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

// errorHandler 限流后返回前端
func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func Rate(status ...string) gin.HandlerFunc {
	var store ratelimit.Store

	if status[0] == "redis" {
		store = ratelimit.RedisStore(&ratelimit.RedisOptions{
			RedisClient: redis.NewClient(&redis.Options{
				Addr:     global.CONFIG.Redis.Host + ":" + global.CONFIG.Redis.Port,
				Password: global.CONFIG.Redis.Password,
				DB:       global.CONFIG.Redis.DB,
			}),
			Rate:  time.Second,
			Limit: global.CONFIG.System.Limit,
		})
	} else {
		store = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
			Rate:  time.Second,
			Limit: global.CONFIG.System.Limit,
		})
	}

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	return mw
}
