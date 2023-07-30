// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAmzStreamSubscribe = "t_amz_stream_subscribe"

// AmzStreamSubscribe mapped from table <t_amz_stream_subscribe>
type AmzStreamSubscribe struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProfileID       int64     `gorm:"column:profile_id;comment:站点id" json:"profile_id"`                                          // 站点id
	ShopName        string    `gorm:"column:shop_name;comment:店铺名称" json:"shop_name"`                                            // 店铺名称
	SellerID        string    `gorm:"column:seller_id;not null;comment:卖家Sellerid" json:"seller_id"`                             // 卖家Sellerid
	MarketplaceID   string    `gorm:"column:marketplace_id;not null;comment:市场id" json:"marketplace_id"`                         // 市场id
	ConfigID        int32     `gorm:"column:config_id;not null;comment:nacos配置中的id" json:"config_id"`                            // nacos配置中的id
	SqsArn          string    `gorm:"column:sqs_arn;comment:SQS arn" json:"sqs_arn"`                                             // SQS arn
	ClientID        string    `gorm:"column:client_id;comment:开发者账号 client_id" json:"client_id"`                                 // 开发者账号 client_id
	ClientSecret    string    `gorm:"column:client_secret;comment:开发者账号 client_secret" json:"client_secret"`                     // 开发者账号 client_secret
	AccessToken     string    `gorm:"column:access_token;comment:access token 3600s失效" json:"access_token"`                      // access token 3600s失效
	AccessTokenTTL  int32     `gorm:"column:access_token_ttl;default:3600;comment:access token 3600s失效" json:"access_token_ttl"` // access token 3600s失效
	AccessTokenTime time.Time `gorm:"column:access_token_time;comment:access token创建时间" json:"access_token_time"`                // access token创建时间
	RefreshToken    string    `gorm:"column:refresh_token;comment:开发者账号的 refresh token" json:"refresh_token"`                    // 开发者账号的 refresh token
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`     // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`     // 更新时间
}

// TableName AmzStreamSubscribe's table name
func (*AmzStreamSubscribe) TableName() string {
	return TableNameAmzStreamSubscribe
}
