package conf

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func getCfg() *ini.File{
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}


func GetMysqlConfig(key string) string{
	cfg := getCfg()
	return cfg.Section("mysql").Key(key).String()
}