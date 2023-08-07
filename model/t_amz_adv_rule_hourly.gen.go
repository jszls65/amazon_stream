// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAmzAdvRuleHourly = "t_amz_adv_rule_hourly"

// AmzAdvRuleHourly mapped from table <t_amz_adv_rule_hourly>
type AmzAdvRuleHourly struct {
	ID                           int64     `gorm:"column:id;primaryKey;comment:主键" json:"id"`                                           // 主键
	RuleID                       int64     `gorm:"column:rule_id;not null;comment:对应规则的id" json:"rule_id"`                              // 对应规则的id
	RuleState                    int32     `gorm:"column:rule_state;not null;default:1;comment:规则状态 0-暂停 1-运行中 2-归档" json:"rule_state"` // 规则状态 0-暂停 1-运行中 2-归档
	CampaignID                   int64     `gorm:"column:campaign_id" json:"campaign_id"`
	AutoRuleType                 int32     `gorm:"column:auto_rule_type;comment:自动化规则 1-分时预算 2-分时调价 3-自动竞价策略 4-自动广告位" json:"auto_rule_type"`           // 自动化规则 1-分时预算 2-分时调价 3-自动竞价策略 4-自动广告位
	Budget                       float64   `gorm:"column:budget;comment:预算" json:"budget"`                                                             // 预算
	StartHour                    string    `gorm:"column:start_hour;comment:开始时间" json:"start_hour"`                                                   // 开始时间
	EndHour                      string    `gorm:"column:end_hour;comment:结束时间" json:"end_hour"`                                                       // 结束时间
	BudgetDistributionPercentage float64   `gorm:"column:budget_distribution_percentage;comment:预算分配比例 0-100" json:"budget_distribution_percentage"`   // 预算分配比例 0-100
	BudgetDistribution           float64   `gorm:"column:budget_distribution;comment:预算分配值" json:"budget_distribution"`                                // 预算分配值
	BudgetAdjustment             float64   `gorm:"column:budget_adjustment;comment:预算调整值" json:"budget_adjustment"`                                    // 预算调整值
	BidAdjustmentFrequency       int32     `gorm:"column:bid_adjustment_frequency;comment:调价频率 0-每天 1-每周 2-工作日 3-周末" json:"bid_adjustment_frequency"`  // 调价频率 0-每天 1-每周 2-工作日 3-周末
	BidAdjustmentDay             string    `gorm:"column:bid_adjustment_day;comment:调价" json:"bid_adjustment_day"`                                     // 调价
	BidAdjustment                string    `gorm:"column:bid_adjustment;comment:竞价调整类型 0,0 (原始竞价, --), 0,1 (原始竞价,上浮)" json:"bid_adjustment"`           // 竞价调整类型 0,0 (原始竞价, --), 0,1 (原始竞价,上浮)
	BidStrategyAdjustment        string    `gorm:"column:bid_strategy_adjustment;comment:竞价策略 10, 200 (调整数值百分比,竞价最大值)" json:"bid_strategy_adjustment"` // 竞价策略 10, 200 (调整数值百分比,竞价最大值)
	Operator                     string    `gorm:"column:operator;comment:操作人" json:"operator"`                                                        // 操作人
	CreateTime                   time.Time `gorm:"column:create_time;comment:开始时间" json:"create_time"`                                                 // 开始时间
	UpdateTime                   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"`              // 修改时间
	BeijingStartHour             string    `gorm:"column:beijing_start_hour;comment:北京开始时间" json:"beijing_start_hour"`                                 // 北京开始时间
	BeijingEndHour               string    `gorm:"column:beijing_end_hour;comment:北京结束时间" json:"beijing_end_hour"`                                     // 北京结束时间
	ProfileID                    int64     `gorm:"column:profile_id" json:"profile_id"`
}

// TableName AmzAdvRuleHourly's table name
func (*AmzAdvRuleHourly) TableName() string {
	return TableNameAmzAdvRuleHourly
}