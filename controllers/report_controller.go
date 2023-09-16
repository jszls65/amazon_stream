package controllers

import (
	"amazon_stream/common"
	"amazon_stream/subfunc"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// 报表
type ReportController struct {
}

func (rc ReportController) Index(c *gin.Context) {
	shopName := c.Query("shopName")
	fmt.Println("shopName: ", shopName)
	if shopName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "shopName不能为空",
		})
		return
	}

	accessToken := subfunc.GenAccessToken(shopName)
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "获取accessToken失败",
		})
		return
	}
	report, code, status := rc.CreateReport(shopName, accessToken)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"status": status,
		"body":   report,
	})
}

func (rc ReportController) CreateReport(shopName string, accessToken string) (string, int, string) {

	shopData := common.GetShopDataMap(shopName)

	httpUrl := "https://advertising-api.amazon.com/v2/hsa/keyword/report"

	bodyStr := `{
    "reportDate": "20239101",
    "creativeType": "all",
    "metrics": "clicks,campaignName,campaignId,adGroupName,adGroupId,adId,impressions,cost,dpv14d,attributedDetailPageViewsClicks14d,attributedOrdersNewToBrand14d,attributedOrdersNewToBrandPercentage14d,attributedOrderRateNewToBrand14d,attributedSalesNewToBrand14d,attributedSalesNewToBrandPercentage14d,attributedUnitsOrderedNewToBrand14d,attributedUnitsOrderedNewToBrandPercentage14d"
}`

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(bodyStr))
	common.HandleError(err)
	req.Header.Add("Content-Type", "application/vnd.MarketingStreamSubscriptions.StreamSubscriptionResource.v1.0+json")
	req.Header.Add("Amazon-Advertising-API-ClientId", shopData.ClientID)
	req.Header.Add("Amazon-Advertising-API-Scope", strconv.FormatInt(shopData.ProfileID, 10))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	fmt.Println("请求头: \n", common.ToJsonStr(req.Header))
	resp, err := http.DefaultClient.Do(req)
	common.HandleError(err)
	defer common.CloseRspBody(resp) // 这步是必要的，防止以后的内存泄漏，切记
	bodyByte, _ := io.ReadAll(resp.Body)

	return string(bodyByte), resp.StatusCode, resp.Status
}
