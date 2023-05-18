package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/grpc-todolist/idl/task/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
	"github.com/CocaineCong/grpc-todolist/pkg/res"
	"github.com/CocaineCong/grpc-todolist/pkg/util"
)

func GetTaskList(ginCtx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskShow(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskCreate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskUpdate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteTask(ginCtx *gin.Context) {
	var tReq pb.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(pb.TaskServiceClient)
	TaskResp, err := TaskService.TaskDelete(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
