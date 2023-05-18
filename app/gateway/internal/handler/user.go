package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	userPb "github.com/CocaineCong/grpc-todolist/idl/user"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
	"github.com/CocaineCong/grpc-todolist/pkg/res"
	"github.com/CocaineCong/grpc-todolist/pkg/util"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq userPb.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(userPb.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq userPb.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(userPb.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Data:   res.TokenData{User: userResp.UserDetail, Token: token},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
