// @Title 查询订阅的状态
// @Author  zls  2023/8/7 22:04
package subfunc

import (
	"amazon_stream/common"
	"log"
	"net/http"
	"strconv"
)

func GetSubState(shopName string, accessToken string, subId string) {
	shopData := common.GetShopDataMap(shopName)
	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions/" + subId
	req, err := http.NewRequest("GET", httpUrl, nil)
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	bodyJsonStr := common.GetRespBodyStr(resp.Body)
	log.Println(bodyJsonStr)
}
