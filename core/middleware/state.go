package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m4n5ter/cnsoftbei/common/log"
	"github.com/m4n5ter/cnsoftbei/server/config"
	"github.com/redis/go-redis/v9"
)

func Redis(conf config.Config) gin.HandlerFunc {
	client := redis.NewClient(&redis.Options{
		Addr:             conf.Redis.Host,
		Password:         conf.Redis.Password,
		DB:               conf.Redis.DB,
		DialTimeout:      time.Duration(conf.Redis.DialTimeout),
		DisableIndentity: conf.Redis.DisableIndentity,
		MaxRetries:       conf.Redis.MaxRetries,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Panic(err)
	}

	return func(c *gin.Context) {
		c.Set("redis", client)
		c.Next()
	}
}
