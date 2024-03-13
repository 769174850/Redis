package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //服务器地址
		Password: "",               //没有密码，默认值
		DB:       0,                //默认DB 0
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil { //链接失败返回错误
		log.Println(err)
		return
	}
	fmt.Println("redis链接成功")
}
