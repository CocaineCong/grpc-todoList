package repository

import (
	"task/pkg/util"
	"os"
)

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&Task{},
		)
	if err != nil {
		util.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	util.LogrusObj.Infoln("register table success")
}