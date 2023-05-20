package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/CocaineCong/grpc-todolist/app/user/internal/repository/db/model"
	userPb "github.com/CocaineCong/grpc-todolist/idl/pb/user"
	"github.com/CocaineCong/grpc-todolist/pkg/util/logger"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// GetUserInfo 获取用户信息
func (dao *UserDao) GetUserInfo(req *userPb.UserRequest) (r *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", req.UserName).
		First(&r).Error

	return
}

// CreateUser 用户创建
func (dao *UserDao) CreateUser(req *userPb.UserRequest) (err error) {
	var user model.User
	var count int64
	dao.Model(&model.User{}).Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}

	user = model.User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	if err = dao.Model(&model.User{}).Create(&user).Error; err != nil {
		logger.LogrusObj.Error("Insert User Error:" + err.Error())
		return
	}

	return
}
