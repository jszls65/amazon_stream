// Package subfunc
// @Author  zls  2023/7/30 11:30
package subfunc

import (
	"amazon_stream/common"
	"log"
	"net/http"
	"strconv"
)

// ListSub 查询店铺的所有订阅
func ListSub(shopName string, accessToken string) {
	shopData := common.GetShopDataMap(shopName)
	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions?maxResults=10"
	req, err := http.NewRequest("GET", httpUrl, nil)
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	bodyJsonStr := common.GetRespBodyStr(resp.Body)
	//log.Println(bodyJsonStr)
	//校验数据集订阅状态
	checkDataSetState(bodyJsonStr)
}

// 校验数据集订阅状态
func checkDataSetState(bodyJsonStr string) {
	bodyMap, err := common.JsonToMap(bodyJsonStr)
	if err != nil {
		log.Fatalln("")
	}
	dataSetSlice := common.GetDataSetSlice()
	// 定义结果map
	resultSubItemMap := make(map[string]string)
	for _, val := range dataSetSlice {
		resultSubItemMap[val] = "--"
	}
	itemSlice := bodyMap.Get("subscriptions").InterSlice()
	itemLen := len(itemSlice)
	if itemSlice == nil || itemLen == 0 {
		for k, v := range resultSubItemMap {
			log.Println(k, ":", v)
		}
		return
	}
	for i := 0; i < itemLen; i++ {
		objxMap := bodyMap.Get("subscriptions[" + strconv.Itoa(i) + "]").MustObjxMap()
		dataSetId := objxMap.Get("dataSetId").String()

		if status, isE := resultSubItemMap[dataSetId]; isE && status == "ACTIVE" {
			continue
		}
		status := objxMap.Get("status").String()
		resultSubItemMap[dataSetId] = status
	}
	// 遍历item
	for k, v := range resultSubItemMap {
		log.Println(k, ":", v)
	}

}
