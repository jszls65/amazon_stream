// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOperateRecord = "t_amz_adv_history_operate_record_sync"

// OperateRecord mapped from table <t_amz_adv_history_operate_record_sync>
type OperateRecord struct {
	ProfileID              int64  `gorm:"column:profile_id;comment:站点id" json:"profile_id"`                                                                                                                                                                                                                               // 站点id
	PortfolioID            int64  `gorm:"column:portfolio_id;comment:组合id" json:"portfolio_id"`                                                                                                                                                                                                                           // 组合id
	CampaignID             int64  `gorm:"column:campaign_id;comment:广告活动id" json:"campaign_id"`                                                                                                                                                                                                                           // 广告活动id
	AdGroupID              int64  `gorm:"column:ad_group_id;comment:广告组id" json:"ad_group_id"`                                                                                                                                                                                                                            // 广告组id
	EntityID               int64  `gorm:"column:entity_id;comment:实体id" json:"entity_id"`                                                                                                                                                                                                                                 // 实体id
	NewValue               string `gorm:"column:new_value;comment:新数据" json:"new_value"`                                                                                                                                                                                                                                  // 新数据
	PreviousValue          string `gorm:"column:previous_value;comment:老数据" json:"previous_value"`                                                                                                                                                                                                                        // 老数据
	EntityType             string `gorm:"column:entity_type;comment:实体类型" json:"entity_type"`                                                                                                                                                                                                                             // 实体类型
	ChangeType             string `gorm:"column:change_type;comment:更改字段" json:"change_type"`                                                                                                                                                                                                                             // 更改字段
	AdAsin                 string `gorm:"column:ad_asin;comment:Field only for ads" json:"ad_asin"`                                                                                                                                                                                                                       // Field only for ads
	PredefinedTarget       string `gorm:"column:predefined_target;comment:SUBSTITUTES, COMPLEMENTS, LOOSE-MATCH, CLOSE-MATCH" json:"predefined_target"`                                                                                                                                                                   // SUBSTITUTES, COMPLEMENTS, LOOSE-MATCH, CLOSE-MATCH
	ProductTargetingType   string `gorm:"column:product_targeting_type;comment:Field only for Product Targeting" json:"product_targeting_type"`                                                                                                                                                                           // Field only for Product Targeting
	TargetingExpression    string `gorm:"column:targeting_expression;comment:Field only for Product Targeting where type = EXPRESSION. Some examples: category="Car Seat Canopies & Covers" brand="Munchkin" price<14.9 rating>4.6 asin="B000NPPATS" category="Women's Handbag Accessories"" json:"targeting_expression"` // Field only for Product Targeting where type = EXPRESSION. Some examples: category="Car Seat Canopies & Covers" brand="Munchkin" price<14.9 rating>4.6 asin="B000NPPATS" category="Women's Handbag Accessories"
	Keyword                string `gorm:"column:keyword;comment:Field only for keywords or negative keywords. The keyword or phrase this matches." json:"keyword"`                                                                                                                                                        // Field only for keywords or negative keywords. The keyword or phrase this matches.
	KeywordType            string `gorm:"column:keyword_type;comment:Field only for keywords" json:"keyword_type"`                                                                                                                                                                                                        // Field only for keywords
	PlacementGroupPosition string `gorm:"column:placement_group_position;comment:Field only for campaign placement group change types" json:"placement_group_position"`                                                                                                                                                   // Field only for campaign placement group change types
	CampaignBudgetType     string `gorm:"column:campaign_budget_type;comment:Field only for campaigns" json:"campaign_budget_type"`                                                                                                                                                                                       // Field only for campaigns
	NegativeTargetingType  string `gorm:"column:negative_targeting_type;comment:Field only for negative targeting." json:"negative_targeting_type"`                                                                                                                                                                       // Field only for negative targeting.
	Status                 int32  `gorm:"column:status;default:1;comment:1-启用 0-暂停" json:"status"`                                                                                                                                                                                                                        // 1-启用 0-暂停
	OperateTime            int64  `gorm:"column:operate_time;comment:操作时间" json:"operate_time"`                                                                                                                                                                                                                           // 操作时间
}

// TableName OperateRecord's table name
func (*OperateRecord) TableName() string {
	return TableNameOperateRecord
}