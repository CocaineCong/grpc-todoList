package dao

import (
	"os"

	"github.com/CocaineCong/grpc-todolist/pkg/util"

	"github.com/CocaineCong/grpc-todolist/app/user/internal/repository/db/model"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		util.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	util.LogrusObj.Infoln("register table success")
}
