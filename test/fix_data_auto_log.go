package main

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"encoding/json"
	"fmt"
)

func main() {
	db := datasource.GetDB()
	// 查询没有广告组合的活动
	var autoLogs []model.AdvAutoLog
	db.Where(" portfolio_id is null").Find(&autoLogs)
	
	// 活动id:广告组合id map
	campaignIds := make([]int64,0)
	for _, autoLog := range autoLogs{
		campaignIds = append(campaignIds, autoLog.CampaignID)
	}
	if len(campaignIds) == 0{
		fmt.Println("活动列表为空")
		return
	}

	// map 活动id:广告组合id
	relMap := make(map[int64]int64, 100)

	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileList := make([]model.AdvCampaigns,0)
	db.Where("campaignId in ? and portfolioid is not null", campaignIds).Find(&campaignPortileList)
	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileListSB := make([]model.AdvCampaignsSb,0)
	db.Where("campaignId in ?  and portfolioid is not null", campaignIds).Find(&campaignPortileListSB)
	// 查询活动表, 构建活动id:广告组合id map
	campaignPortileListSD := make([]model.AdvCampaignsSd,0)
	db.Where("campaignId in ?  and portfolioid is not null", campaignIds).Find(&campaignPortileListSD)

	if len(campaignPortileList) > 0{
		for _,val := range campaignPortileList{
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	if len(campaignPortileListSB) > 0{
		for _,val := range campaignPortileListSB{
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	if len(campaignPortileListSD) > 0{
		for _,val := range campaignPortileListSD{
			relMap[val.CampaignID] = val.Portfolioid
		}
	}

	fmt.Println("relMap: ", relMap)

	for _, val := range autoLogs{
		portfolioId, ok :=relMap[val.CampaignID]
		if !ok{
			continue
		}
		val.PortfolioID = portfolioId
	}

	jsonv, _ :=json.Marshal(relMap)
	fmt.Println("relMap:", string(jsonv))

	for cid,pid := range relMap{

		// sql :=fmt.Sprintf("update t_amz_adv_auto_log set portfolio_id = %d where campaign_id = %d", pid,cid)
		// fmt.Println("sql=", sql)

		err := db.Exec("update t_amz_adv_auto_log set portfolio_id = ? where campaign_id = ?", pid,cid).Error
		if err != nil {
			fmt.Println("执行sql失败, ", err.Error())
			return 
		}
	}
}

