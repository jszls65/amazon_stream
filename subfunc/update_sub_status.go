package subfunc

import (
	"amazon_stream/common"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 更新订阅状态
func UpdateSubStatus(shopName string, accessToken string, subId string, status string) {
	shopData := common.GetShopDataMap(shopName)
	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions/" + subId
	bodyMap := make(map[string]string)
	bodyMap["notes"] = "广告" + common.GetRandomToken(10)
	bodyMap["status"] = status

	fmt.Println("入参: ", common.ToJsonStr(bodyMap))

	bodyStr, err := common.MapToJson(bodyMap)
	common.HandleError(err)

	req, err := http.NewRequest("PUT", httpUrl, strings.NewReader(bodyStr))
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
