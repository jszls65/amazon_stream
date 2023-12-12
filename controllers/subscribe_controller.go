package controllers

import (
	"amazon_stream/datasource"
	"amazon_stream/model"
	"amazon_stream/service/subfunc"
	"database/sql"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type SubscribeController struct {
}

func (this SubscribeController) GetInfo(c *gin.Context) {
	shopName := c.Query("shopName")
	if shopName == "" {
		c.JSON(400, gin.H{
			"data": "参数异常",
		})
		return
	}
	subInfo := getSubscribeResults(shopName)
	c.JSON(200, gin.H{
		"data": subInfo,
	})
}

func (this SubscribeController) Create(c *gin.Context) {
	shopName := c.Query("shopName")
	dataSetId := c.Query("dataSetId") // sd-traffic"
	// 参数校验
	if shopName == "" || dataSetId == "" {
		c.JSON(400, gin.H{
			"data": "参数异常",
		})
		return
	}
	split := strings.Split(dataSetId, ",")
	var wg sync.WaitGroup
	wg.Add(len(split))
	resultList := make([]string, 0)

	// 创建token
	accessToken := subfunc.GenAccessToken(shopName)
	for _, ds := range split {
		go func(ds string) {
			// 创建订阅
			resp := subfunc.CreateSub(shopName, accessToken, ds)
			if resp.StatusCode == 200 {
				resultList = append(resultList, ds+" 订阅成功")
			} else {
				resultList = append(resultList, ds+" 订阅失败:"+resp.Status)
			}
			wg.Done()
		}(ds)
	}
	wg.Wait()
	c.JSON(200, gin.H{
		"msg":  "执行结束",
		"data": resultList,
	})
}

// 查询订阅结果
func getSubscribeResults(shopName string) []string {
	accessToken := subfunc.GenAccessToken(shopName)
	return subfunc.ListSub(shopName, accessToken)
}

// 生成stream表数据
func (this SubscribeController) GenStreamData(context *gin.Context) {

	defer func() {
		msg := recover()
		if msg != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": msg})
			return
		}
	}()

	shopName, ok := context.GetQuery("shopName")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数异常",
		})
		return
	}
	db := datasource.GetDB()

	// 查询店铺market表
	market := this.getShopMarket(shopName)
	sellerId := market.Sellerid
	marketplaceId := market.MarketplaceID

	// 查询auth表数据
	auth := this.getAmazonAuth(sellerId)

	// 查询订阅表, 根据configId来获取clientId和Secret
	ss := this.getStreamSubInfo(auth.ConfigID)
	// 查询profile
	profile := this.getProfile(sellerId, marketplaceId)
	streamSub := &model.StreamSubscribe{
		ShopName:        shopName,
		ProfileID:       profile,
		SellerID:        sellerId,
		MarketplaceID:   marketplaceId,
		ConfigID:        auth.ConfigID,     // auth表
		RefreshToken:    auth.RefreshToken, // auth表
		ClientID:        ss.ClientID,
		ClientSecret:    ss.ClientSecret,
		AccessTokenTime: time.Now(),
	}
	// 保存订阅表数据
	err := db.Save(streamSub).Error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "执行成功"})
}

func (this SubscribeController) getAmazonAuth(sellerId string) model.AmazonAuth {
	var auth model.AmazonAuth
	datasource.GetDB().Raw(`
		select * from t_amazon_auth where sellerid = @sellerId limit 1
	`, sql.Named("sellerId", sellerId)).Scan(&auth)
	if auth.Sellerid == "" {
		panic("t_amazon_auth 无数据")
	}
	return auth
}

func (this SubscribeController) getShopMarket(shopName string) model.AmazonsellerMarket {
	var market model.AmazonsellerMarket
	//datasource.GetDB().Raw(`
	//	select * from t_amazonseller_market where name = @name limit 1
	//`, sql.Named("name", shopName)).Scan(&market)
	datasource.GetDB().Model(model.AmazonsellerMarket{Name: shopName}).Find(&market)

	if market.Sellerid == "" {
		panic("t_amazonseller_market 无数据")
	}
	return market
}

func (this SubscribeController) getStreamSubInfo(configId int32) model.StreamSubscribe {
	var ss model.StreamSubscribe
	//datasource.GetDB().Raw(`
	//	select * from t_amz_stream_subscribe where config_id = @configId limit 1
	//`, sql.Named("configId", configId)).Scan(&ss)
	err := datasource.GetDB().Model(model.StreamSubscribe{ConfigID: configId}).Find(&ss).Error
	if err != nil {
		panic(err.Error())
	}
	return ss
}

func (this SubscribeController) getProfile(sellerId string, marketplaceId string) int64 {
	var profile model.AdvProfile
	err := datasource.GetDB().Model(model.AdvProfile{SellerID: sellerId, MarketplaceID: marketplaceId}).Find(&profile).Error
	if err != nil {
		panic("getProfile查询sql失败: " + err.Error())
	}
	return profile.ID
}
