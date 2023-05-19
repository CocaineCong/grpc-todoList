package handler

import (
	"context"
	"sync"

	"github.com/CocaineCong/grpc-todolist/app/user/internal/repository/db/dao"
	userPb "github.com/CocaineCong/grpc-todolist/idl/user/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
	userPb.UnimplementedUserServiceServer
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	resp = new(userPb.UserDetailResponse)
	resp.Code = e.SUCCESS
	r, err := dao.NewUserDao(ctx).GetUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.UserDetail = &userPb.UserResponse{
		UserId:   r.UserID,
		UserName: r.UserName,
		NickName: r.UserName,
	}
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.CommonResponse, err error) {
	resp.Code = e.SUCCESS
	err = dao.NewUserDao(ctx).CreateUser(req)
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	resp.Data = e.GetMsg(uint(resp.Code))
	return
}

func (u *UserSrv) UserLogout(ctx context.Context, request *userPb.UserRequest) (resp *userPb.CommonResponse, err error) {
	return
}
