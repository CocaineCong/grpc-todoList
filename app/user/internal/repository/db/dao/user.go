package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/CocaineCong/grpc-todolist/idl/user"
	"github.com/CocaineCong/grpc-todolist/pkg/util"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (user *User) CheckUserExist(req *user.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

func (*User) Create(req *service.UserRequest) error {
	var user User
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exist")
	}
	user = User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	if err := DB.Create(&user).Error; err != nil {
		util.LogrusObj.Error("Insert User Error:" + err.Error())
		return err
	}
	return nil
}

// 视图返回
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserID),
		NickName: item.NickName,
		UserName: item.UserName,
	}
	return &userModel
}
