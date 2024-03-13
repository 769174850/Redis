package content

import (
	"fmt"
	"redis/dao"
)

func Content() {
	for {
		fmt.Println("欢迎来到排行榜系统，选择不同数字进入不同功能")
		fmt.Println("数字1获取排行榜总人数和排行榜数据\n" + "数字2搜索用户及其分数和排名\n" + "数字3添加用户及其分数\n" + "数字4删除排行榜内的用户\n" + "数字5更改用户的分数\n" + "数字6退出系统")
		var num int
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Println("请选择正序查询倒序查询（数字1为正序查询，数字2为倒序查询")
			var number int
			fmt.Scanln(&number)
			if number == 1 {
				fmt.Println("执行获取排行榜正序数据操作...")
				fmt.Println("执行获取排行榜总人数和排行榜数据操作...")
				dao.GetRank()       //获取排行榜正序数据
				dao.GetRankNumber() //获取排行榜总人数和排行榜数据
			} else {
				fmt.Println("执行获取排行榜倒序数据操作...")
				fmt.Println("执行获取排行榜总人数和排行榜数据操作...")
				dao.GetRankDesc()   //获取排行榜倒序数据
				dao.GetRankNumber() //获取排行榜总人数和排行榜数据
			}

		case 2:
			fmt.Println("执行检索用户所输入的数据操作...")
			dao.Research() //进行检索用户所输入的数据
		case 3:
			fmt.Println("执行添加用户及其分数操作...")
			dao.Add() //添加用户及其分数
		case 4:
			fmt.Println("执行删除用户操作...")
			dao.Delete() //删除用户
		case 5:
			fmt.Println("执行更改用户分数操作...")
			dao.Update() //更改用户分数
		case 6:
			fmt.Println("退出系统")
			return
		}
		fmt.Println("是否返回主界面(Y/N)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			fmt.Println("返回主界面...")
		} else {
			fmt.Println("退出程序")
			return
		}
	}
}
