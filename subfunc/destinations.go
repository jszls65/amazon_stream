package subfunc

import (
	"amazon_stream/common"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CreateDestinations 创建destinations
func CreateDestinations(shopName string) *http.Response {
	log.Println("创建订阅方法入参,shopName:", shopName)
	shopData := common.GetShopDataMap(shopName)
	sqsArn := shopData.SqsArn
	if sqsArn == "" {
		panic("店铺[" + shopName + "]的sqs arn为空, 请补充订阅表中的数据")
	}
	// url
	httpUrl := "https://sellingpartnerapi-na.amazon.com/notifications/v1/destinations"
	// body
	bodyMapStr := `{
	  "name": "{name}",
	  "resourceSpecification":
	  {
		"sqs":
		{
		  "arn": "{sqs}"
		}
	  }
	}`
	bodyMapStr = strings.ReplaceAll(bodyMapStr, "{name}", getDestName(sqsArn))
	bodyMapStr = strings.ReplaceAll(bodyMapStr, "{sqs}", sqsArn)
	fmt.Println("入参: ", bodyMapStr)
	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(bodyMapStr))
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Amazon-Advertising-API-withScopes", "SCOPE_NOTIFICATIONS_API")
	//req.Header.Add("Authorization", "Bearer "+GenAccessToken(shopName))
	fmt.Println("请求头: \n", common.ToJsonStr(req.Header))
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	defer common.CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记
	fmt.Println("返回code:", resp.StatusCode, " msg:", resp.Status)
	return resp
}

func getDestName(arn string) string {
	split := strings.Split(arn, ":")
	splitLen := len(split)
	if splitLen < 2 {
		return common.GetId("dest_name_")
	}

	name := "dest_name_" + split[splitLen-2]
	fmt.Println("destinationName:", name)
	return name

}
