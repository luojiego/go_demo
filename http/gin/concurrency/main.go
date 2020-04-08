package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	num = 0
)

//用来测试web环境 并发对于单个变量 操作的表现情况
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		num++
		c.String(http.StatusOK, strconv.FormatInt(int64(num), 10))
	})
	r.Run(":7001")
}
