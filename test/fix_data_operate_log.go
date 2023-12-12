package main

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func main() {
	db := datasource.GetDB()
	limit := 500
	offset := 0
	for {
		if !doRun2(db, limit, offset) {
			break
		}
		fmt.Println("----------offset:", strconv.FormatInt(int64(offset), 10), "-----------")
		offset += limit
	}
}

// 执行处理, 分页处理
func doRun2(db *gorm.DB, limit int, offset int) bool {
	// 查询业务表
	var operateLogs []model.OperateLog
	db.Where("campaignId is not null and campaignId != 0 and portfolio_id is null").Limit(limit).Offset(offset).Find(&operateLogs)
	fmt.Println("当前业务表数据量:", len(operateLogs))
	if len(operateLogs) == 0 {
		return false // 没有数据,表明处理结束
	}

	// 收集业务数据中的活动id切片
	campaignIds := make([]int64, 0)
	for _, item := range operateLogs {
		currCid := item.CampaignID
		if currCid == 0 {
			continue
		}
		campaignIds = append(campaignIds, currCid)
	}
	if len(campaignIds) == 0 {
		fmt.Println("活动列表为空")
		return false
	}
	fmt.Println("需要处理的活动数量:", len(campaignIds))

	relMap := getRelMap(db, campaignIds)
	fmt.Println("relMap的容量:", len(relMap))

	if len(relMap) == 0 {
		return true
	}

	for _, val := range operateLogs {
		portfolioId, ok := relMap[val.CampaignID]
		if !ok {
			continue
		}
		val.PortfolioID = portfolioId
	}

	// 开始更新业务表
	for cid, pid := range relMap {
		err := db.Exec("update t_amz_adv_operate_log set portfolio_id = ? where campaignId = ? and portfolio_id is null",
			pid, cid).Error

		if err != nil {
			fmt.Println("执行sql失败, ", err.Error())
			return false
		}
	}
	return true
}

// map 活动id:广告组合id 关系
func getRelMap(db *gorm.DB, campaignIds []int64) map[int64]int64 {
	// map 活动id:广告组合id
	relMap := make(map[int64]int64, 100)

	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileList := make([]model.AdvCampaigns, 0)
	db.Where("campaignId in ? and portfolioid is not null", campaignIds).Find(&campaignPortileList)
	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileListSB := make([]model.AdvCampaignsSb, 0)
	db.Where("campaignId in ?  and portfolioid is not null", campaignIds).Find(&campaignPortileListSB)
	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileListSD := make([]model.AdvCampaignsSd, 0)
	db.Where("campaignId in ?  and portfolioid is not null", campaignIds).Find(&campaignPortileListSD)

	if len(campaignPortileList) > 0 {
		for _, val := range campaignPortileList {
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	if len(campaignPortileListSB) > 0 {
		for _, val := range campaignPortileListSB {
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	if len(campaignPortileListSD) > 0 {
		for _, val := range campaignPortileListSD {
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	return relMap
}
