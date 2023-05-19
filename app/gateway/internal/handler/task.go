package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/grpc-todolist/idl/task/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
	"github.com/CocaineCong/grpc-todolist/pkg/res"
	"github.com/CocaineCong/grpc-todolist/pkg/util/ctl"
)

func GetTaskList(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	TaskService := ctx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskShow(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func CreateTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	TaskService := ctx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskCreate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func UpdateTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	TaskService := ctx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskUpdate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func DeleteTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	TaskService := ctx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskDelete(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}
