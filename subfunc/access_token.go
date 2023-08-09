// @Title
// @Author  zls  2023/7/30 11:45
package subfunc

import (
	"amazon_stream/common"
	"amazon_stream/pojo"
	"encoding/json"
	"io"
	"net/http"
)

// GenAccessToken 生成 access token
// shopIndex 店铺号
func GenAccessToken(shopName string) string {

	shopData := common.GetShopDataMap(shopName)
	httpUrl := "https://api.amazon.com/auth/o2/token?client_id=" + shopData.ClientID + "&client_secret=" + shopData.ClientSecret + "&grant_type=refresh_token&refresh_token=" + shopData.RefreshToken
	req, _ := http.NewRequest("POST", httpUrl, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	defer common.CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记

	body, err := io.ReadAll(resp.Body) // 读取响应 body, 返回为 []byte
	common.HandleError(err)
	//log.Println("请求成功, body:", string(body))

	// 将json字符串映射到结构体中
	var rsp pojo.AccessTokenResp
	err = json.Unmarshal(body, &rsp)
	common.HandleError(err)
	return rsp.AccessToken
}
