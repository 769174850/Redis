package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func GetRankNumber() {
	number, err := rdb.ZCard(context.Background(), "rank").Result() //使用ZCard函数获取排行榜人数
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("获取排行榜数量失败")
		return
	}

	fmt.Println("总人数:", number) //输出排行榜人数
}

func GetRank() {
	var number2 int64
	fmt.Println("请输入自己要查的名次（例如：查到第三名输入3）：注：后一个数字输入0即为查询剩余全部")
	fmt.Scanln(&number2)
	ranking, err := rdb.ZRevRangeWithScores(context.Background(), "rank", 0, number2-1).Result() //使用ZRevRangeWithScores函数获取排行榜数据
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("获取排行榜失败")
		return
	}

	fmt.Println("排行榜:")
	for _, z := range ranking { //将ZRevRangeWithScores函数获取的数组一一输出
		fmt.Println(z.Member, z.Score)
	}
}

func GetRankDesc() {
	var number2 int64
	fmt.Println("请输入自己要查的人数（例如：1 3）：注：后一个数字输入0即为查询剩余全部")
	fmt.Scanln(&number2)
	ranking, err := rdb.ZRangeWithScores(context.Background(), "rank", 0, number2-1).Result() //使用ZRangeWithScores函数获取排行榜数据
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("获取排行榜失败")
		return
	}

	fmt.Println("排行榜:")
	for _, z := range ranking { //将ZRangeWithScores函数获取的数组一一输出
		fmt.Println(z.Member, z.Score)
	}
}

func Research() {
	fmt.Println("请输入你想要查找的成员:")
	var member string
	fmt.Scanln(&member)                                                     //输入成员姓名
	Score, err := rdb.ZScore(context.Background(), "rank", member).Result() //使用ZRank函数获取成员信息
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("这个成员不存在! ")
		return
	}
	rank, err := rdb.ZRevRank(context.Background(), "rank", member).Result() //使用ZRevRank函数获取成员排名
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("这个成员不存在! ")
		return
	}
	fmt.Println(member, "的分数是", Score, "他/她的排名为", rank+1)
}

func Add() {
	fmt.Println("请输入你要添加的成员及其分数: ")
	var person string
	var score float64
	fmt.Println("成员姓名:")
	fmt.Scanln(&person) //输入成员姓名
	fmt.Println("分数:")
	fmt.Scanln(&score) //输入成员分数

	exist, err := rdb.ZRank(context.Background(), "rank", person).Result()
	if err != nil {
		log.Println(err) //返回错误日志
		return
	}

	if exist >= 0 {
		fmt.Println("这个成员已经存在! ")
		return
	}

	err = rdb.ZAdd(context.Background(), "rank", redis.Z{Score: score, Member: person}).Err() //使用ZAdd函数添加成员信息
	if err != nil {
		log.Println(err) //返回错误日志
		return
	}

	fmt.Println("成员添加成功!")
}

func Delete() {
	fmt.Println("请输入你想删除的成员:")
	var person string
	fmt.Scanln(&person)

	err := rdb.ZRem(context.Background(), "rank", person)
	if err != nil {
		log.Println(err) //返回错误日志
		fmt.Println("这个成员不存在! ")
		return
	}

	fmt.Println(person, "已经成功删除了!")
}

func Update() {
	fmt.Println("请输入你想要更新的成员姓名及其分数: ")
	var person string
	var score float64
	fmt.Println("成员姓名:")
	fmt.Scanln(&person) //输入成员姓名
	fmt.Println("分数:")
	fmt.Scanln(&score) //输入成员分数

	err := rdb.ZAdd(context.Background(), "rank", redis.Z{Score: score, Member: person}) //使用ZAdd函数添加成员信息
	if err != nil {
		log.Println(err) //返回错误日志
		return
	}

	fmt.Println(person, "的分数已经被更新为", score)
}
