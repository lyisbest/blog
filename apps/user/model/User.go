package model

type User struct {
	Id       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Ctime    int64  `gorm:"column:ctime"`
	Mtime    int64  `gorm:"column:mtime"`
}

func (User) TableName() string {
	return "user_tab"
}

type UserLog struct {
	Id            int    `gorm:"column:id" json:"id"`
	Operator      string `gorm:"column:operator" json:"operator"`
	OperationType int    `gorm:"column:operation_type" json:"operation_type"`
	Operation     string `gorm:"column:operation" json:"operation"`
	Time          int64  `gorm:"column:time" json:"time"`
}

func (UserLog) Table() string {
	return "user_log_tab"
}
