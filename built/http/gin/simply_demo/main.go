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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "luojie"})
	})
	r.GET("/test", func(c *gin.Context) {
		if err := Test(); err != nil {
			c.String(http.StatusOK, "Error accours\n")
			return
		}
		c.String(http.StatusOK, "Good! ")
	})
	r.Static("/.well-known", "./data")
	r.Run(":8080")
}

func Test() error {
	return nil
}
