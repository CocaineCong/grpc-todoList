package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context,req *service.UserRequest) (resp *service.UserDetailResponse,err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse,err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.Create(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp,err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp,nil
}

func (*UserService) UserLogout (ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse,err error) {
	resp = new(service.UserDetailResponse)
	return resp, nil
}