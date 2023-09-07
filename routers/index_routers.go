package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRouterInit(r *gin.Engine) {
	r.GET("index", controllers.IndexController{}.Index)
}
