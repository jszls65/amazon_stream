package controllers

import (
	"amazon_stream/service/subfunc"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type SubscribeController struct {
}

func (sub SubscribeController) GetInfo(c *gin.Context) {
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
}

func (sub SubscribeController) Create(c *gin.Context) {
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
	accessToken := subfunc.GenAccessToken(shopName)
	for _, ds := range split {
		go func(ds string) {
			// 创建订阅
			resp := subfunc.CreateSub(shopName, accessToken, ds)
			if resp.StatusCode == 200 {
				resultList = append(resultList, ds+" 订阅成功")
			} else {
				resultList = append(resultList, ds+" 订阅失败:"+resp.Status)
			}
			wg.Done()
		}(ds)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"msg":  "执行结束",
		"data": resultList,
	})
}

// 查询订阅结果
func getSubscribeResults(shopName string) []string {
	accessToken := subfunc.GenAccessToken(shopName)
	return subfunc.ListSub(shopName, accessToken)
}
