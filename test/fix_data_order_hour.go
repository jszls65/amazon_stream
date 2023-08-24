// @Title 测试, 修复表中转换后的北京时间错误问题
// @Author  zls  2023/8/4 13:21
package main

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"fmt"
	"log"
	"strings"
	"time"
)

// 修复订单小时级别表的历史数据
func main() {
	db := datasource.GetDB()
	var salesEvents []model.AmsSalesEvent
	db.Raw("select * from t_ams_sales_event where seller_id is null or seller_id = '' or start_hour is null   ").Scan(&salesEvents)
	total := len(salesEvents)
	log.Println("返回总数据量: ", total)
	if total == 0 {
		log.Println("无数据")
		return
	}
	for _, val := range salesEvents {
		// 时间-小时
		startHourInt := getHourInt32(val.StartTime)
		// 店铺id
		sellerId := getSellerId(val.AccountID)
		fmt.Printf("id: %d, startHour: %d, sellerId: %s \n", val.ID, startHourInt, sellerId)
		db.Model(model.AmsSalesEvent{ID: val.ID}).Updates(model.AmsSalesEvent{
			StartHour: startHourInt,
			SellerID:  sellerId,
		})
		//val.StartHour = startHourInt
		//val.SellerID = sellerId
	}
	//fmt.Println(salesEvents)
	// 批量更新
	//db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"start_hour", "seller_id"}),
	//}).Create(&salesEvents)
}

func getSellerId(id string) string {
	if id == "" {
		return ""
	}
	split := strings.Split(id, ".")
	return split[len(split)-1]
}

func getHourInt32(startTime time.Time) int32 {
	return int32(startTime.Hour())
}
