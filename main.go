// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/subfunc"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
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
		split := strings.Split(dataSetId, ",")
		var wg sync.WaitGroup
		wg.Add(len(split))
		resultList := make([]string, 0)

		// 创建token
		accessToken = subfunc.GenAccessToken(shopName)
		for _, ds := range split {
			go func(ds string) {
				// 创建订阅
				resp := subfunc.CreateSub(shopName, accessToken, ds)
				if resp.StatusCode == 200 {
					resultList = append(resultList, ds+" 订阅成功")
				} else {
					resultList = append(resultList, ds+" 订阅失败")
				}
				wg.Done()
			}(ds)
		}
		wg.Wait()
		c.JSON(200, gin.H{
			"msg":  "执行结束",
			"data": resultList,
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

// 查询订阅结果
func getSubscribeResults(shopName string) []string {
	accessToken = subfunc.GenAccessToken(shopName)
	return subfunc.ListSub(shopName, accessToken)
}
