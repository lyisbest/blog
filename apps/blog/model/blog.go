package model

type Blog struct {
	Id      int    `gorm:"column:id" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Creator string `gorm:"column:creator" json:"creator"`
	Ctime   int64  `gorm:"column:ctime" json:"ctime"`
	Mtime   int64  `gorm:"column:mtime" json:"mtime"`
}

func (Blog) TableName() string {
	return "blog_tab"
}

type BlogView struct {
	Id      int   `gorm:"column:id" json:"id"`
	ViewNum int64 `gorm:"view_num" json:"view_num"`
}

func (BlogView) TableName() string {
	return "blog_view_tab"
}

type BlogLog struct {
	Id            int    `gorm:"column:id" json:"id"`
	BlogId        int    `gorm:"column:blog_id" json:"blogId"`
	Operator      string `gorm:"column:operator" json:"operator"`
	OperationType int    `gorm:"column:operation_type" json:"operation_type"`
	Operation     string `gorm:"column:operation" json:"operation"`
	Ctime         int64  `gorm:"column:ctime" json:"ctime"`
}

func (BlogLog) TableName() string {
	return "blog_log_tab"
}
