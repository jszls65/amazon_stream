// @Title
// @Author  zls  2023/7/28 22:15
package common

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
)

func GetShopDataMap(shopName string) (shopData *model.AmzStreamSubscribe) {
	db := datasource.GetDB()
	var amzStreamSubscribe model.AmzStreamSubscribe
	db.Raw("select * from t_amz_stream_subscribe where shop_name = ?", shopName).Scan(&amzStreamSubscribe)
	return &amzStreamSubscribe
}

func GetShopName(index string) string {
	var ShopNameMap = make(map[string]string, 0)
	ShopNameMap["33"] = "美亚三十三"
	ShopNameMap["1"] = "美亚一"
	ShopNameMap["36"] = "美亚三十六"
	return ShopNameMap[index]
}
