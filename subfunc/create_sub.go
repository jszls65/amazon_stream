// @Title
// @Author  zls  2023/7/30 11:47
package subfunc

import (
	"amazon_stream/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CreateSub 创建订阅
func CreateSub(shopName string, accessToken string, dataSetId string) *http.Response {
	log.Println("创建订阅方法入参,shopName:", shopName, "; dataSetId:", dataSetId)
	shopData := common.GetShopDataMap(shopName)

	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions"

	bodyMap := make(map[string]string)
	bodyMap["notes"] = "广告" + common.GetRandomToken(10)
	bodyMap["clientRequestToken"] = common.GetRandomToken(25)
	bodyMap["dataSetId"] = dataSetId
	bodyMap["destinationArn"] = shopData.SqsArn

	fmt.Println("入参: ", common.ToJsonStr(bodyMap))

	bodyStr, err := common.MapToJson(bodyMap)
	common.HandleError(err)

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(bodyStr))
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	fmt.Println("请求头: \n", common.ToJsonStr(req.Header))
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	defer common.CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记
	fmt.Println("返回code:", resp.StatusCode, " msg:", resp.Status)
	respBodyBytes, err := io.ReadAll(resp.Body) // 读取响应 body, 返回为 []byte
	common.HandleError(err)
	respBodyStr := string(respBodyBytes)
	fmt.Println("请求结束, body:", respBodyStr)
	if !strings.Contains(respBodyStr, "subscriptionId") {
		fmt.Println("创建订阅失败")
	}
	return resp
}
