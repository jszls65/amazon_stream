// @Title 测试, 修复表中转换后的北京时间错误问题
// @Author  zls  2023/8/4 13:21
package main

import (
	"amazon_stream/common"
	"amazon_stream/datasource"
	"amazon_stream/model"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	db := datasource.GetDB()
	var hourlies []model.AmzAdvRuleHourly
	db.Raw("select * from t_amz_adv_rule_hourly where beijing_start_hour is not null").Scan(&hourlies)
	total := len(hourlies)
	log.Println("返回总数据量: ", total)
	if total == 0 {
		log.Println("无数据")
		return
	}
	for _, val := range hourlies {
		startHourInt := getHourInt(val.StartHour)
		endHourInt := getHourInt(val.EndHour)
		if checkDataOk(val) {
			continue
		}
		bjStartHourInt := startHourInt + 15
		if bjStartHourInt >= 24 {
			bjStartHourInt -= 24
		}
		bjEndHourInt := bjStartHourInt + (endHourInt - startHourInt)
		if bjEndHourInt > 24 {
			bjEndHourInt -= 24
		}
		bjStartHourStr := getHourStr(bjStartHourInt)
		bjEndHourStr := getHourStr(bjEndHourInt)
		db.Model(model.AmzAdvRuleHourly{ID: val.ID}).Updates(model.AmzAdvRuleHourly{
			BeijingStartHour: bjStartHourStr,
			BeijingEndHour:   bjEndHourStr,
		})
		log.Println("更新前: ID:", val.ID, "\t BeijingStartHour:", val.BeijingStartHour,
			"\t BeijingEndHour:", val.BeijingEndHour)
		log.Println("更新后: ID:", val.ID, "\t BeijingStartHour:", bjStartHourStr,
			"\t BeijingEndHour:", bjEndHourStr)
		time.Sleep(time.Millisecond * 200) // 睡眠200毫秒
	}
}

// 校验数据是否正确, 正确-true  错误-false
func checkDataOk(hourly model.AmzAdvRuleHourly) bool {
	startHourInt := getHourInt(hourly.StartHour)
	endHourInt := getHourInt(hourly.EndHour)
	bjStartHourInt := getHourInt(hourly.BeijingStartHour)
	bjEndHourInt := getHourInt(hourly.BeijingEndHour)
	if bjStartHourInt >= 24 || bjEndHourInt > 24 {
		return false
	}
	if bjStartHourInt < startHourInt {
		bjStartHourInt += 24
	}
	if bjStartHourInt-startHourInt != 15 {
		return false
	}

	if bjEndHourInt < endHourInt {
		bjEndHourInt += 24
	}
	if bjEndHourInt-endHourInt != 15 {
		return false
	}

	return true
}

func getHourInt(hourStr string) int {
	startHourSplit := strings.Split(hourStr, ":")
	atoi, err := strconv.Atoi(startHourSplit[0])
	common.HandleError(err)
	return atoi
}

func getHourStr(hourInt int) string {
	if hourInt < 10 {
		return "0" + strconv.Itoa(hourInt) + ":00"
	}
	return strconv.Itoa(hourInt) + ":00"
}
