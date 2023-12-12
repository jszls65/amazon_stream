// @Title 测试, 修复表中转换后的北京时间错误问题
// @Author  zls  2023/8/4 13:21
package main

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

// 修复订单小时级别表的历史数据
func main2() {
	db := datasource.GetDB()
	for {
		hasData := updateCampaignId4Log(db, 100)
		if !hasData {
			break
		}
	}
}

func updateCampaignId4Log(db *gorm.DB, size int) bool {
	var operateLogs []model.AmzAdvOperateLog
	db.Where("beanclasz = ? and campaignId is null", "AmzAdvProductTargeSD").Limit(size).Find(&operateLogs)
	total := len(operateLogs)
	log.Println("返回总数据量: ", total)
	if total == 0 {
		log.Println("无数据")
		return false
	}
	// 查询广告组表
	// 广告组id列表
	advGroupIds := make([]int64, 0)
	for _, val := range operateLogs {
		advGroupIds = append(advGroupIds, val.AdGroupID)
	}
	// 获取广告组id - 活动id map
	groupIdCampaignIdMap := getGroupIdCampaignIdMap(advGroupIds)

	for _, val := range operateLogs {
		val.CampaignID = groupIdCampaignIdMap[val.AdGroupID]
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"campaignId"}),
	}).Create(&operateLogs)

	return true
}

func getGroupIdCampaignIdMap(advGroupIds []int64) map[int64]int64 {
	db := datasource.GetDB()
	var adgroupList []model.AmzAdvAdgroupsSd
	db.Where("adGroupId in ?", advGroupIds).Find(&adgroupList)
	// 返回值
	ret := make(map[int64]int64, 0)
	for _, val := range adgroupList {
		ret[val.AdGroupID] = val.CampaignID
	}
	fmt.Println(ret)
	return ret
}
