package controllers

import (
	"amazon_stream/service/subfunc"

	"github.com/gin-gonic/gin"
)

type PolicyController struct {
}

// 生成策略
func (sub PolicyController) GenPolicy(c *gin.Context) {
	shopName := c.Query("shopName")
	// 参数校验
	if shopName == "" {
		c.JSON(400, gin.H{
			"data": "参数异常",
		})
		return
	}
	policy := subfunc.GenSqsPolicy(shopName)
	c.JSON(200, gin.H{
		"msg":  "执行成功",
		"data": policy,
	})
}
