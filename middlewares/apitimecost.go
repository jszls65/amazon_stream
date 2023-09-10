// @Title
// @Author  zls  2023/9/7 22:30
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// CostApiTime 计算接口耗时 时长
func CostApiTime(c *gin.Context) {
	startTime := time.Now().UnixMilli()

	c.Next()

	endTime := time.Now().UnixMilli()
	fmt.Println("接口:", c.FullPath(), "  耗时: ", float64(endTime-startTime)/1000, "s")
}
