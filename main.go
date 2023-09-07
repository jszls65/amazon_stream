// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载html
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	routers.InitRouters(r)

	err := r.Run(":8082")
	common.HandleError(err)
}
