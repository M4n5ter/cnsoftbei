package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m4n5ter/cnsoftbei/common/log"
	"github.com/m4n5ter/cnsoftbei/server/config"
	"github.com/redis/go-redis/v9"
)

func Redis(conf config.Config) gin.HandlerFunc {
	client := redis.NewClient(&redis.Options{
		Addr:             fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password:         conf.Redis.Password,
		DB:               conf.Redis.DB,
		DialTimeout:      time.Second * time.Duration(conf.Redis.DialTimeout),
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