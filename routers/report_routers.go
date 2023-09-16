package routers

import (
	"amazon_stream/controllers"

	"github.com/gin-gonic/gin"
)

func ReportRouterInit(r *gin.Engine) {
	r.GET("/report", controllers.ReportController{}.Index)
}
