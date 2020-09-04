/**
* @Author : henry
* @Data: 2020-08-17 20:22
* @Note:
**/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建默认路由
	r := gin.Default()
	//r.Use(AuthMiddleware())
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"127.0.0.1", false, true)
		// 返回信息
		c.String(http.StatusOK, "Login success")
	})

	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	r.Run(":8000")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "err",
		})
		// 若验证不通过,不再调用后续的函数处理
		c.Abort()
		return
	}
}
