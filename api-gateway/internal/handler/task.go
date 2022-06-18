package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
)

func GetTaskList(ginCtx *gin.Context) {
	var fReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&fReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	fReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskShow(context.Background(), &fReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var fReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&fReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	fReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskCreate(context.Background(), &fReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var fReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&fReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	fReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskUpdate(context.Background(), &fReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}


func DeleteTask(ginCtx *gin.Context) {
	var fReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&fReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	fReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskDelete(context.Background(), &fReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
