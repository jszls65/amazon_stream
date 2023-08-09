// @Title
// @Author  zls  2023/7/28 22:18
package common

var dataSetSlice = []string{
	"sp-traffic",
	"sp-conversion",
	"budget-usage",
	"sd-traffic",
	"sd-conversion",
	"campaigns",
	"adgroups",
	"ads",
	"targets",
	"sponsored-ads-campaign-diagnostics-recommendations",
}

func GetDataSetSlice() []string {
	return dataSetSlice
}

// TimeTemplate 时间模板
const TimeTemplate string = "2006-01-02 15:04:05"
