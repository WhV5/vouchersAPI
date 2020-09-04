/**
* @Author : henry
* @Data: 2020-08-13 15:40
* @Note:
**/

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vouchersAPI/app"
	"net/http"
)

type Voucher interface {
	AddVoucher()
	SelectVoucher() Voucher
}

var logger = app.Logger

type kisUV struct{}

func AddVoucher(c *gin.Context) {
	// 绑定参数

	c.JSON(http.StatusOK, "Hello")
}
