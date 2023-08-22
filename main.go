// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/subfunc"
	"github.com/gin-gonic/gin"
	"net/http"
)

var accessToken string

func main() {
	r := gin.Default()
	// 查询订阅信息
	r.GET("subscribe/getInfo", func(c *gin.Context) {
		shopName := c.Query("shopName")
		if shopName == "" {
			c.JSON(400, gin.H{
				"data": "参数异常",
			})
			return
		}
		subInfo := getSubscribeResults(shopName)
		c.JSON(200, gin.H{
			"data": subInfo,
		})
	})
	// 创建订阅信息
	r.GET("subscribe/create", func(c *gin.Context) {
		shopName := c.Query("shopName")
		dataSetId := c.Query("dataSetId") // sd-traffic"
		// 参数校验
		if shopName == "" || dataSetId == "" {
			c.JSON(400, gin.H{
				"data": "参数异常",
			})
			return
		}
		// 创建订阅
		response := create(shopName, dataSetId)
		c.JSON(200, gin.H{
			"msg":  "执行成功",
			"data": response,
		})
	})

	// 生成配置
	r.GET("subscribe/genPolicy", func(c *gin.Context) {
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
	})

	// 创建sqs destination
	r.GET("subscribe/destinations", func(c *gin.Context) {
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
	})

	err := r.Run(":8082")
	common.HandleError(err)
}

// 创建订阅
func create(shopName string, dataSetId string) *http.Response {
	if accessToken == "" {
		accessToken = subfunc.GenAccessToken(shopName)
	}
	return subfunc.CreateSub(shopName, accessToken, dataSetId)
}

// 查询订阅结果
func getSubscribeResults(shopName string) []string {
	if accessToken == "" {
		accessToken = subfunc.GenAccessToken(shopName)
	}
	return subfunc.ListSub(shopName, accessToken)
}
