// @Title
// @Author  zls  2023/7/27 08:27
package main

import (
	"amazon_stream/common"
	"amazon_stream/middlewares"
	"amazon_stream/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 中间件
	r.Use(middlewares.CostApiTime)
	// 加载html
	r.LoadHTMLGlob("templates/**/*")
	// 参数一: 路由,  参数二: 静态文件根目录
	r.Static("/static", "./static")

	routers.InitRouters(r)

	err := r.Run(":8082")
	common.HandleError(err)
}
