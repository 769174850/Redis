package main

import (
	"redis/content"
	"redis/dao"
)

func main() {
	dao.Init()
	content.Content()
}
