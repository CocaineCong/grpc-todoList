package service

import (
	"context"
	"sync"

	"github.com/CocaineCong/grpc-todolist/app/task/internal/repository/db/dao"
	pb "github.com/CocaineCong/grpc-todolist/idl/pb/task"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
	pb.UnimplementedTaskServiceServer
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}
func (*TaskSrv) TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).CreateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(int(resp.Code))
	return
}

func (*TaskSrv) TaskShow(ctx context.Context, req *pb.TaskRequest) (resp *pb.TasksDetailResponse, err error) {
	resp = new(pb.TasksDetailResponse)
	r, err := dao.NewTaskDao(ctx).ListTaskByUserId(req.UserID)
	resp.Code = e.SUCCESS
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	for i := range r {
		resp.TaskDetail = append(resp.TaskDetail, &pb.TaskModel{
			TaskID:    r[i].TaskID,
			UserID:    r[i].UserID,
			Status:    int64(r[i].Status),
			Title:     r[i].Title,
			Content:   r[i].Content,
			StartTime: r[i].StartTime,
			EndTime:   r[i].EndTime,
		})
	}
	return
}

func (*TaskSrv) TaskUpdate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(int(resp.Code))
	return
}

func (*TaskSrv) TaskDelete(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).DeleteTaskById(req.TaskID, req.UserID)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(int(resp.Code))
	return
}
