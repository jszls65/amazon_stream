// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAdvCampaigns = "t_amz_adv_campaigns"

// AdvCampaigns mapped from table <t_amz_adv_campaigns>
type AdvCampaigns struct {
	CampaignID           int64     `gorm:"column:campaignId;primaryKey" json:"campaignId"`
	Profileid            int64     `gorm:"column:profileid;not null" json:"profileid"`
	Portfolioid          int64     `gorm:"column:portfolioid;comment:广告组合 id" json:"portfolioid"`        // 广告组合 id
	Name                 string    `gorm:"column:name;comment:广告活动名称" json:"name"`                       // 广告活动名称
	CampaignType         string    `gorm:"column:campaignType;comment:sp 和 sb（hsa）" json:"campaignType"` // sp 和 sb（hsa）
	DailyBudget          float64   `gorm:"column:dailyBudget;comment:每日预算" json:"dailyBudget"`           // 每日预算
	TargetingType        string    `gorm:"column:targetingType;comment:投放类型" json:"targetingType"`       // 投放类型
	State                string    `gorm:"column:state" json:"state"`
	PremiumBidAdjustment string    `gorm:"column:premiumBidAdjustment;comment:竞价" json:"premiumBidAdjustment"` // 竞价
	Bidding              string    `gorm:"column:bidding" json:"bidding"`
	StartDate            time.Time `gorm:"column:startDate;comment:开始时间" json:"startDate"` // 开始时间
	EndDate              time.Time `gorm:"column:endDate" json:"endDate"`
	Opttime              time.Time `gorm:"column:opttime" json:"opttime"`
}

// TableName AdvCampaigns's table name
func (*AdvCampaigns) TableName() string {
	return TableNameAdvCampaigns
}
