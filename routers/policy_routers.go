package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func PolicyRoutersInit(r *gin.Engine) {

	g := r.Group("policy")
	{
		// 查询订阅信息
		g.GET("genPolicy", controllers.PolicyController{}.GenPolicy)
	}

}
