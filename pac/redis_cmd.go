package pac

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

// RedisCmd 执行 Redis 命令
func RedisCmd(cmd string) interface{} {

	ctx := context.Background()

	var argsInterface []interface{}

	// 处理输入字符串有空格的问题
	if strings.Contains(cmd, "\"") {
		oldString := ReString(cmd, "\"(.*?)\"")
		newString := strings.ReplaceAll(oldString, " ", "$")
		cmd = strings.ReplaceAll(cmd, oldString, newString)
		cmd = strings.ReplaceAll(cmd, "\"", "")
	}

	args := strings.Fields(cmd)
	for _, arg := range args {
		if strings.Contains(arg, "$") {
			arg = strings.ReplaceAll(arg, "$", " ")
		}
		argsInterface = append(argsInterface, arg)
	}

	info, err := Rdb.Do(ctx, argsInterface...).Result()
	if err != nil {
		Err(err)
		return ""
	}
	return info
}

// 获取 Redis 基本信息
func redisVersion() bool {
	info := RedisCmd("info")
	if strings.Contains(info.(string), "redis_version") {
		Info("获取 Redis 基本信息")
		os := ReString(info, "os:.*")
		version := ReString(info, "redis_version:.*")
		Success(os)
		Success(version)
		dir := RedisCmd("config get dir")
		redisDir = redisString(dir)[4:]
		Success(redisDir)

		file := RedisCmd("config get dbfilename")
		redisDbFilename = redisString(file)[11:]
		Success(redisDbFilename)
		return true
	}
	return false
}

// 循环执行shell命令
func loopCmd(s string) {
	Info("执行命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			if strings.Contains(s, "exec") {
				CloseSlave("exec")
			}
			break
		}
		// 执行命令
		if strings.Contains(s, "exec") {
			RunCmd(cmd)
		} else if strings.Contains(s, "lua") {
			RedisLua(cmd)
		}

	}
}

// 循环执行 Redis 命令
func loopRedis() {
	Info("执行 Redis 命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		// 执行命令
		fmt.Println(RedisCmd(cmd))
	}
}
