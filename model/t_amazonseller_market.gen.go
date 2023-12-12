// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAmazonsellerMarket = "t_amazonseller_market"

// AmazonsellerMarket mapped from table <t_amazonseller_market>
type AmazonsellerMarket struct {
	Sellerid        string    `gorm:"column:sellerid;primaryKey;comment:卖家Sellerid" json:"sellerid"`       // 卖家Sellerid
	MarketplaceID   string    `gorm:"column:marketplace_id;primaryKey;comment:站点ID" json:"marketplace_id"` // 站点ID
	MarketPlaceName string    `gorm:"column:market_place_name" json:"market_place_name"`
	Country         string    `gorm:"column:country;comment:国家编码" json:"country"`                        // 国家编码
	Name            string    `gorm:"column:name;comment:站点英文名称" json:"name"`                            // 站点英文名称
	Language        string    `gorm:"column:language;comment:对应语言编码" json:"language"`                    // 对应语言编码
	Currency        string    `gorm:"column:currency;comment:对应币种" json:"currency"`                      // 对应币种
	Domain          string    `gorm:"column:domain;comment:对应域名" json:"domain"`                          // 对应域名
	Amazonauthid    int64     `gorm:"column:amazonauthid;comment:授权对应ID等同于Sellerid" json:"amazonauthid"` // 授权对应ID等同于Sellerid
	Opttime         time.Time `gorm:"column:opttime;comment:操作时间" json:"opttime"`                        // 操作时间
	Disable         []uint8   `gorm:"column:disable;default:b'0;comment:操作人" json:"disable"`             // 操作人
}

// TableName AmazonsellerMarket's table name
func (*AmazonsellerMarket) TableName() string {
	return TableNameAmazonsellerMarket
}
