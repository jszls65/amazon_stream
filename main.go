// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/subfunc"
)

func main() {
	// 生成token
	accessToken := subfunc.GenAccessToken(common.ShopIndex36)
	// 创建所有类型的订阅
	//dataSetSlice := common.GetDataSetSlice()
	//subfunc.CreateSub(common.ShopIndex36, accessToken, "sp-traffic")
	//for _, val := range dataSetSlice {
	//	log.Println(val)
	//	//	subfunc.CreateSub(common.ShopIndex36, accessToken, val)
	//}
	subfunc.ListSub(common.ShopIndex36, accessToken)
}
