package rpc

import (
	"context"
	"errors"

	taskPb "github.com/CocaineCong/grpc-todolist/idl/pb/task"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

func TaskCreate(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskCreate(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskUpdate(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskUpdate(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskDelete(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TaskCommonResponse, err error) {
	r, err := TaskClient.TaskDelete(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New(r.Msg)
		return
	}

	return r, nil
}

func TaskList(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TasksDetailResponse, err error) {
	r, err := TaskClient.TaskShow(ctx, req)

	if err != nil {
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.New("获取失败")
		return
	}

	return r, nil
}
