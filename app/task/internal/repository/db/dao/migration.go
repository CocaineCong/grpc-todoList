package dao

import (
	"os"

	"github.com/CocaineCong/grpc-todolist/app/task/internal/repository/db/model"
	"github.com/CocaineCong/grpc-todolist/pkg/util"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Task{},
		)
	if err != nil {
		util.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	util.LogrusObj.Infoln("register table success")
}
