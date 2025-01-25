package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 密码，如果没有密码则为空字符串
		DB:       0,                // 使用的数据库编号
	})

	// Ping测试连接
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	cmd := client.Dump(context.Background(), `hash_test`)
	str := cmd.Val()

	parsedData, err := parseCustomDump([]byte(str))
	if err != nil {
		fmt.Println("Error parsing Redis hash:", err)
		return
	}

	// 输出解析结果
	for key, value := range parsedData {
		fmt.Printf("%s: %s\n", key, value)
	}
}

// parseCustomDump 解析 Redis DUMP 数据（自定义格式）
func parseCustomDump(data []byte) (map[string]string, error) {
	result := make(map[string]string)
	buffer := bytes.NewBuffer(data)

	// length, _ := readLength(buffer)
	// bs := make([]byte, length)
	// buffer.Read(bs)
	// fmt.Println(string(bs))

	// 开始解析键值对
	buffer.Next(12)
	for buffer.Len() > 8 { // 剩余数据大于校验和长度
		// 读取键的长度
		key, _ := parseStr(buffer)
		value, _ := parseStr(buffer)
		// 保存键值对
		result[string(key)] = string(value)
	}

	// 忽略校验和部分
	return result, nil
}

// readLength 读取长度
// return length, error

// func parseKey(buffer *bytes.Buffer) (string, error) {
// 	length, err := buffer.ReadByte()
// 	if err != nil {
// 		return "", fmt.Errorf("")
// 	}
// }

// return string, usually for key parse
