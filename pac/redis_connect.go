package pac

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var (
	Rdb             *redis.Client
	redisDir        string
	redisDbFilename string
)

// RedisClient 连接 Redis
func RedisClient(pwd string) (err error) {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", Rhost, Rport),
		Password: pwd, // 密码认证
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	if strings.Contains(pong, "PONG") {
		redisVersion()
		//fmt.Println("连接成功")
	}
	return nil
}
