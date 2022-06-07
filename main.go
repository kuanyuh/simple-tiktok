package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/dao"
)

func main() {
	r := gin.Default()
	dao.Init() //初始化gorm
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
