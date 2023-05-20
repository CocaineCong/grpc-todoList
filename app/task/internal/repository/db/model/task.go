package model

type Task struct {
	TaskID    int64 `gorm:"primarykey"` // id
	UserID    int64 `gorm:"index"`      // 用户id
	Status    int   `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (*Task) Table() string {
	return "task"
}
