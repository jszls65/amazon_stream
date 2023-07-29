// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	. "amazon_stream/pojo"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 生成token
	accessToken := genAccessToken(ShopIndex36)
	// 创建
	//createSub(ShopIndex36, accessToken, "")
	listSub(ShopIndex36, accessToken)
}

// 生成 access token
// shopIndex 店铺号
func genAccessToken(shopIndex int) string {

	shopData := getShopDataMap(shopIndex)
	httpUrl := "https://api.amazon.com/auth/o2/token?client_id=" + shopData.ClientId + "&client_secret=" + shopData.ClientSecret + "&grant_type=refresh_token&refresh_token=" + shopData.RefreshToken
	req, _ := http.NewRequest("POST", httpUrl, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	HandleError(err)
	defer CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记

	body, err := io.ReadAll(resp.Body) // 读取响应 body, 返回为 []byte
	HandleError(err)
	log.Println("请求成功, body:", string(body))

	// 将json字符串映射到结构体中
	var rsp AccessTokenResp
	err = json.Unmarshal(body, &rsp)
	HandleError(err)
	return rsp.AccessToken
}

// 创建订阅
func createSub(shopIndex int, accessToken string, dataSetId string) {
	log.Println("创建订阅方法入参,shopIndex:", shopIndex, "; dataSetId:", dataSetId)
	shopData := getShopDataMap(shopIndex)

	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions"

	bodyMap := make(map[string]string)
	bodyMap["notes"] = "广告"
	bodyMap["clientRequestToken"] = "广告" + strconv.FormatInt(time.Now().Unix(), 10)
	bodyMap["dataSetId"] = dataSetId

	bodyStr, err := MapToJson(bodyMap)
	HandleError(err)

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(bodyStr))
	HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientId)
	req.Header.Add("Amazon-Advertising-API-Scope", shopData.ProfileId)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	HandleError(err)
	defer CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记

	respBodyBytes, err := io.ReadAll(resp.Body) // 读取响应 body, 返回为 []byte
	HandleError(err)
	respBodyStr := string(respBodyBytes)
	log.Println("请求成功, body:", respBodyStr)

	if !strings.Contains(respBodyStr, "subscriptionId") {
		panic("创建订阅失败")
	}
}

// 查询店铺的所有订阅
func listSub(shopIndex int, accessToken string) {
	shopData := getShopDataMap(shopIndex)
	httpUrl := "https://advertising-api.amazon.com/streams/subscriptions?maxResults=10"
	req, err := http.NewRequest("GET", httpUrl, nil)
	HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientId)
	req.Header.Add("Amazon-Advertising-API-Scope", shopData.ProfileId)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	HandleError(err)
	log.Println(GetRespBodyStr(resp.Body))
}
