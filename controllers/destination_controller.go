package controllers

import (
	"amazon_stream/service/subfunc"

	"github.com/gin-gonic/gin"
)

type DestinationController struct{}

func (d DestinationController) Create(c *gin.Context) {
	shopName := c.Query("shopName")
	// 参数校验
	if shopName == "" {
		c.JSON(400, gin.H{
			"data": "参数异常",
		})
		return
	}
	// 发送http请求
	destinations := subfunc.CreateDestinations(shopName)
	c.JSON(200, gin.H{
		"msg":  "执行成功",
		"data": destinations,
	})
}
