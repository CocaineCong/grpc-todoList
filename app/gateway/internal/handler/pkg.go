package handler

import (
	"errors"

	"github.com/CocaineCong/grpc-todolist/pkg/util"
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
