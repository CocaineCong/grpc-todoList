package repository

import (
	"task/internal/service"
	"task/pkg/util"
)

type Task struct {
	TaskID    uint `gorm:"primarykey"` // id
	UserID    uint `gorm:"index"`      // 用户id
	Status    int  `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (*Task) Show(req *service.TaskRequest) (taskList []Task, err error) {
	err = DB.Model(Task{}).Where("user_id=?", req.UserID).Find(&taskList).Error
	if err != nil {
		return taskList, err
	}
	return taskList, nil
}

func (*Task) Create(req *service.TaskRequest) error {
	task := Task{
		UserID:    uint(req.UserID),
		Title:     req.Title,
		Content:   req.Content,
		Status:    int(req.Status),
		StartTime: int64(req.StartTime),
		EndTime:   int64(req.EndTime),
	}
	if err := DB.Create(&task).Error; err != nil {
		util.LogrusObj.Error("Insert Task Error:" + err.Error())
		return err
	}
	return nil
}

func (*Task) Delete(req *service.TaskRequest) error {
	err := DB.Where("task_id=?", req.TaskID).Delete(Task{}).Error
	return err
}

func (*Task) Update(req *service.TaskRequest) error {
	t := Task{}
	err := DB.Where("task_id=?", req.TaskID).First(&t).Error
	if err != nil {
		return err
	}
	t.Title = req.Title
	t.Content = req.Content
	t.Status = int(req.Status)
	t.StartTime = int64(req.StartTime)
	t.EndTime = int64(req.EndTime)
	err = DB.Save(&t).Error
	return err
}

func BuildTasks(item []Task) (tList []*service.TaskModel) {
	for _, v := range item {
		f := BuildTask(v)
		tList = append(tList, f)
	}
	return tList
}

func BuildTask(item Task) *service.TaskModel {
	return &service.TaskModel{
		TaskID:    uint32(item.TaskID),
		UserID:    uint32(item.UserID),
		Status:    uint32(item.Status),
		Title:     item.Title,
		Content:   item.Content,
		StartTime: uint32(item.StartTime),
		EndTime:   uint32(item.EndTime),
	}
}
