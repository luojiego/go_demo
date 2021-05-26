package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//用来测试服务器的防火墙是否配置OK
//最简单的办法：使用python开启一个web服务
//python -m SimpleHTTPServer 8080
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "<h1>hello</h1>")
	})
	r.Static("/.well-known", "./data")
	r.Run(":8080")
}
