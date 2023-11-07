// @Title
// @Author  zls  2023/9/7 20:32
package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {

	IndexRouterInit(r)
	PolicyRoutersInit(r)
	SubscribeRoutersInit(r)
	DestinationsRouterInit(r)
	ReportRouterInit(r)
	SqsRouterInit(r)
}
