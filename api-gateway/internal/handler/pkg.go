package handler

import (
	"api-gateway/pkg/util"
	"errors"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}
