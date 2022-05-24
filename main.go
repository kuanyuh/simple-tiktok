package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/dao"
)

func main() {
	dao.Init()
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
