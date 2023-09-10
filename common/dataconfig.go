// @Title
// @Author  zls  2023/7/28 22:15
package common

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"time"
)

func GetShopDataMap(shopName string) (shopData *model.AmzStreamSubscribe) {
	db := datasource.GetDB()
	var amzStreamSubscribe model.AmzStreamSubscribe
	db.Raw("select * from t_amz_stream_subscribe where shop_name = ?", shopName).Scan(&amzStreamSubscribe)
	return &amzStreamSubscribe
}

func SaveAccessToken(token string, id int64) {

	db := datasource.GetDB()
	amzStreamSubscribe := &model.AmzStreamSubscribe{
		ID:              id,
		AccessToken:     token,
		AccessTokenTime: time.Now(),
		AccessTokenTTL:  3600,
	}
	db.Model(&amzStreamSubscribe).Updates(amzStreamSubscribe)
}
