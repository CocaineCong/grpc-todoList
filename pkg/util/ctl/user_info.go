package ctl

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/grpc-todolist/consts"
)

type UserInfo struct {
	Id int64 `json:"id"`
}

func GetUserInfo(ctx *gin.Context) (*UserInfo, error) {
	return &UserInfo{Id: ctx.GetInt64(consts.UserIdKey)}, nil
}

func InitUserInfo(ctx context.Context) {
	// TOOD 放缓存，之后的用户信息，走缓存
}
