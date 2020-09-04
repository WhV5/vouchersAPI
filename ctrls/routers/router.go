/**
* @Author : henry
* @Data: 2020-08-13 15:07
* @Note: 初始化 gin
**/

package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	r.GET("/kis/voucher/add", AddVoucher)

	return r
}
