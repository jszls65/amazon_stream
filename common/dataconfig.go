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
