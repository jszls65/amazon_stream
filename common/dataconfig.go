// @Title
// @Author  zls  2023/7/28 22:15
package common

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"time"
)

func GetShopDataMap(shopName string) (shopData *model.StreamSubscribe) {
	db := datasource.GetDB()
	var StreamSubscribe model.StreamSubscribe
	db.Raw("select * from t_amz_stream_subscribe where shop_name = ?", shopName).Scan(&StreamSubscribe)
	return &StreamSubscribe
}

func SaveAccessToken(token string, id int64) {

	db := datasource.GetDB()
	StreamSubscribe := &model.StreamSubscribe{
		ID:              id,
		AccessToken:     token,
		AccessTokenTime: time.Now(),
		AccessTokenTTL:  3600,
	}
	db.Model(&StreamSubscribe).Updates(StreamSubscribe)
}
