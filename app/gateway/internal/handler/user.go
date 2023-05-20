package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/CocaineCong/grpc-todolist/idl/pb/user"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
	"github.com/CocaineCong/grpc-todolist/pkg/res"
	"github.com/CocaineCong/grpc-todolist/pkg/util/jwt"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userReq pb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ctx.Keys["user"].(pb.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userReq pb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ctx.Keys["user"].(pb.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := jwt.GenerateToken(userResp.UserDetail.UserId)
	r := res.Response{
		Data:   res.TokenData{User: userResp.UserDetail, Token: token},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}
