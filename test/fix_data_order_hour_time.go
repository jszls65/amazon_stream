// @Title 测试, 修复表中转换后的北京时间错误问题
// @Author  zls  2023/8/4 13:21
package main

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"fmt"
	"log"
	"time"
)

// 修复订单小时级别表的历史数据

func main() {
	db := datasource.GetDB()
	var salesEvents []model.AmsSalesEvent
	db.Raw("select * from t_ams_sales_event where profile_start_time is null and marketplace_id = 'ATVPDKIKX0DER' ").Scan(&salesEvents)
	total := len(salesEvents)
	log.Println("返回总数据量: ", total)
	if total == 0 {
		log.Println("无数据")
		return
	}
	for i, val := range salesEvents {
		// 时间-小时
		profileStartTime := val.StartTime.Add(-7 * time.Hour)
		profileEndTime := val.EndTime.Add(-7 * time.Hour)
		db.Model(model.AmsSalesEvent{ID: val.ID}).Updates(model.AmsSalesEvent{
			ProfileStartTime: profileStartTime,
			ProfileEndTime:   profileEndTime,
		})
		fmt.Println("已处理", (i+1))
		time.Sleep(100 * time.Millisecond)

	}
	fmt.Println("执行结束")

}
