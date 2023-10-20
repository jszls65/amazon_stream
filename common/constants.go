// @Title
// @Author  zls  2023/7/28 22:18
package common

var dataSetSlice = []string{
	"sp-traffic",
	"sp-conversion",
	"budget-usage",
	"sd-traffic",
	"sd-conversion",
	"sb-traffic",
	"sb-conversion",
	"sb-clickstream",
	"campaigns",
	"adgroups",
	"ads",
	"targets",
	"sponsored-ads-campaign-diagnostics-recommendations",
}

func GetDataSetSlice() []string {
	return dataSetSlice
}
