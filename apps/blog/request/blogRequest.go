package request

import (
	"bytes"
)

type BlogCreateRequest struct {
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
}

type BlogUpdateRequest struct {
	Id      int    `gorm:"column:id" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
}

type BlogDeleteRequest struct {
	Id int `gorm:"column:id" json:"id"`
}

type BlogGetRequest struct {
	Id int `gorm:"column:id" json:"id"`
}

type BlogLogGetRequest struct {
	Id int `gorm:"column:id" json:"id"`
}

type ListBlogRequest struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type ImageRequest struct {
	File bytes.Buffer `json:"file"`
}
