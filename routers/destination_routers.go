package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func DestinationsRouterInit(r *gin.Engine) {
	g := r.Group("destinations")

	g.GET("create", controllers.DestinationController{}.Create)

}
