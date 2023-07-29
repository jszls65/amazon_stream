// @Title
// @Author  zls  2023/7/28 22:15
package main

import (
	. "amazon_stream/pojo"
)

func getShopDataMap(index int) (shopData ShopInfo) {
	info, exist := shopInfoMap[index]
	if exist {
		return info
	}
	shopInfoMap[ShopIndex36] = s36
	return shopInfoMap[index]
}

// 店铺数据map key:店铺号
var shopInfoMap = make(map[int]ShopInfo)

// 美亚36 店铺信息
var s36 = ShopInfo{
	ProfileId:    "**",
	ClientId:     "**",
	ClientSecret: "**",
	RefreshToken: "**",
}
