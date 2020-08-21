package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	num = 0
)

//用来测试web环境 并发 安全性
//结论：无法保证安全性 虽然出现的概率并不会太高 但是一定会出现
//误区：千万不要使用单核机器进行测试
//google cloud：单核表现正常
//aws ec2: 8核出 panic
//win 10: 也会 panic
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		num++
		c.String(http.StatusOK, strconv.FormatInt(int64(num), 10))
	})
	r.Run(":7001")
}
