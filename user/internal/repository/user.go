package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"
	"user/pkg/util"
)

type User struct {
	UserID         uint      `gorm:"primarykey"`
	UserName       string    `gorm:"unique"`
	NickName       string
	PasswordDigest string
}


const (
	PassWordCost = 12 // 密码加密难度
)

// 加密密码
func (user *User) SetPassword (password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error;
		err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) ShowUserInfo (req *service.UserRequest)(err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

func (*User) Create (req *service.UserRequest) error {
	var user User
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count!=0 {
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
