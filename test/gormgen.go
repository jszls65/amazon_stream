// @Title
// @Author  zls  2023/7/29 21:02
package main

import (
	"amazon_stream/datasource"

	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb := datasource.GetDB()
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.GenerateModelAs("t_amz_adv_operate_log", "AmzAdvOperateLog")
	// g.GenerateModelAs("t_amz_adv_campaigns", "AdvCampaigns")
	// g.GenerateModelAs("t_amz_adv_campaigns_sd", "AdvCampaignsSd")
	g.GenerateModelAs("t_amz_adv_profile", "AdvProfile")

	// Generate the code
	// 单独执行命令: go run gormgen.go
	g.Execute()
}
