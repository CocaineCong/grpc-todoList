package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/grpc-todolist/app/task/internal/repository/db/model"
	taskPb "github.com/CocaineCong/grpc-todolist/idl/task/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/util/logger"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) ListTaskByUserId(userId int64) (r []*model.Task, err error) {
	err = dao.Model(&model.Task{}).
		Where("user_id=?", userId).
		Find(&r).Error

	return
}

func (dao *TaskDao) CreateTask(req *taskPb.TaskRequest) (err error) {
	t := &model.Task{
		UserID:    req.UserID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    int(req.Status),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	if err = dao.Model(&model.Task{}).Create(&t).Error; err != nil {
		logger.LogrusObj.Error("Insert Task Error:" + err.Error())
		return
	}

	return
}

func (dao *TaskDao) DeleteTaskById(taskId, userId int64) (err error) {
	err = dao.Model(&model.Task{}).
		Where("task_id = ? AND user_id = ?", taskId, userId).
		Delete(model.Task{}).Error

	return
}

func (dao *TaskDao) UpdateTask(req *taskPb.TaskRequest) (err error) {
	t := model.Task{}
	err = dao.Model(&model.Task{}).
		Where("task_id=?", req.TaskID).First(&t).Error
	if err != nil {
		return
	}
	t.Title = req.Title
	t.Content = req.Content
	t.Status = int(req.Status)
	t.StartTime = req.StartTime
	t.EndTime = req.EndTime
	err = dao.Save(&t).Error

	return
}
