package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func SqsRouterInit(r *gin.Engine) {
	g := r.Group("sqs")
	g.GET("pull-msg", controllers.SqsController{}.PullMsg)

}
