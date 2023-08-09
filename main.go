// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/subfunc"

	"github.com/gin-gonic/gin"
)

var accessToken string
var pathRoot = "subscribe/"

func main() {
	r := gin.Default()
	// 查询订阅信息
	r.GET(pathRoot+"/getSubscribeInfo", func(c *gin.Context) {
		shopIndex := c.Query("shopIndex")

		subInfo := getSubscribeResults(common.GetShopName(shopIndex))
		c.JSON(200, gin.H{
			"data": subInfo,
		})
	})
	// 创建订阅信息
	r.GET(pathRoot+"/create", func(c *gin.Context) {
		shopIndex := c.Query("shopIndex")

		subInfo := getSubscribeResults(common.GetShopName(shopIndex))
		c.JSON(200, gin.H{
			"data": subInfo,
		})
	})
	r.Run(":8082")
}

// 创建订阅
func subscribe(shopName string, index int) {
	if accessToken == "" {
		accessToken = subfunc.GenAccessToken(shopName)
	}
	dataSetSlice := common.GetDataSetSlice()
	accessToken := subfunc.GenAccessToken(shopName)
	subfunc.CreateSub(shopName, accessToken, dataSetSlice[index])
}

// 查询订阅结果
func getSubscribeResults(shopName string) map[string]interface{} {
	if accessToken == "" {
		accessToken = subfunc.GenAccessToken(shopName)
	}
	return subfunc.ListSub(shopName, accessToken)
}
