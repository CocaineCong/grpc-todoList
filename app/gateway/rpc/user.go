package rpc

import (
	"context"
	"errors"

	userPb "github.com/CocaineCong/grpc-todolist/idl/pb/user"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

func UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserResponse, err error) {
	r, err := UserClient.UserLogin(ctx, req)
	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("登陆失败")
		return
	}

	return r.UserDetail, nil
}

func UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserRegister(ctx, req)
	if err != nil {
		return
	}

	if resp.Code != e.SUCCESS {
		err = errors.New(resp.Msg)
		return
	}

	return
}
