package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func SubscribeRoutersInit(r *gin.Engine) {
	g := r.Group("subscribe")
	// 查询订阅信息
	g.GET("/getInfo", controllers.SubscribeController{}.GetInfo)
	g.GET("/create", controllers.SubscribeController{}.Create)
	g.GET("/genStreamData", controllers.SubscribeController{}.GenStreamData)
}
