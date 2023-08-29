// Package subfunc
// @Author  zls  2023/7/30 11:30
package subfunc

import (
	"amazon_stream/common"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ListSub 查询店铺的所有订阅
func ListSub(shopName string, accessToken string) []string {
	if accessToken == "" {
		log.Fatalln("accessToken不能为空")
		return nil
	}
	shopData := common.GetShopDataMap(shopName)
	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions?maxResults=30"
	req, err := http.NewRequest("GET", httpUrl, nil)
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	bodyJsonStr := common.GetRespBodyStr(resp.Body)
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(bodyJsonStr), "", "    ")
	fmt.Println(str.String())
	//校验数据集订阅状态
	return checkDataSetState(bodyJsonStr, shopData.SqsArn)
}

// 校验数据集订阅状态
func checkDataSetState(bodyJsonStr string, sqsArn string) []string {
	bodyMap, err := common.JsonToMap(bodyJsonStr)
	common.HandleError(err)
	dataSetSlice := common.GetDataSetSlice()
	// 定义结果map
	resultSubItemMap := make(map[string]string)
	for _, dataSetId := range dataSetSlice {
		resultSubItemMap[dataSetId] = "--"
	}
	itemSlice := bodyMap.Get("subscriptions").InterSlice()
	itemLen := len(itemSlice)

	for i := 0; i < itemLen; i++ {
		objxMap := bodyMap.Get("subscriptions[" + strconv.Itoa(i) + "]").MustObjxMap()
		dataSetId := objxMap.Get("dataSetId").String()
		destinationArn := objxMap.Get("destinationArn").String()
		if sqsArn != destinationArn {
			continue
		}

		if status, isE := resultSubItemMap[dataSetId]; isE && status == "ACTIVE" {
			continue
		}
		status := objxMap.Get("status").String()
		resultSubItemMap[dataSetId] = status
	}

	resultList := []string{}
	for i, val := range dataSetSlice {
		v := resultSubItemMap[val]
		fmt.Println(i, " ", val, ":", v)
		itemStr := strconv.Itoa(i) + " " + val + ":" + v
		resultList = append(resultList, itemStr)
	}
	return resultList
}
