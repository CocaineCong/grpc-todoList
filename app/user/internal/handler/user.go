package handler

import (
	"context"
	"sync"

	"github.com/CocaineCong/grpc-todolist/app/user/internal/repository/db/dao"
	userPb "github.com/CocaineCong/grpc-todolist/idl/user"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
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
		return resp, err
	}
	userPb.UserModel{
		UserID:   r.UserID,
		UserName: r.UserName,
		NickName: r.UserName,
	}
	resp.UserDetail =
	return resp, nil
}

func (u *UserSrv) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.Create(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (u *UserSrv) UserLogout(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	return resp, nil
}
