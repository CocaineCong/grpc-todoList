package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/CocaineCong/grpc-todolist/idl/pb/task"
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
	taskService := ctx.Keys["task"].(pb.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func CreateTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	taskService := ctx.Keys["task"].(pb.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func UpdateTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	taskService := ctx.Keys["task"].(pb.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func DeleteTask(ctx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ctx.Bind(&tReq))
	user, err := ctl.GetUserInfo(ctx)
	PanicIfTaskError(err)
	tReq.UserID = user.Id
	taskService := ctx.Keys["task"].(pb.TaskServiceClient)
	taskResp, err := taskService.TaskDelete(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}
