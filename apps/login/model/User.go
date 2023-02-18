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
