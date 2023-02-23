package repository

import (
	"blog/apps/blog/model"
	"blog/configuration"
	"errors"
	"golang.org/x/net/context"
	"log"
)

type BlogRepository struct {
}

func NewBlogRepository() *BlogRepository {
	return &BlogRepository{}
}

func (r *BlogRepository) CreateBlog(ctx context.Context, blog *model.Blog) (*model.Blog, error) {
	result := configuration.DB.Create(blog)
	if err := result.Error; err != nil {
		log.Printf("Create Blog failed, blog: %v, error: %v\n", *blog, err)
		return nil, err
	}
	if rowNum := result.RowsAffected; rowNum == 0 {
		err := errors.New("create Blog failed, rowNum: 0")
		log.Printf("Create Blog Failed, blog: %v, error: %v", *blog, err)
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) UpdateBlog(ctx context.Context, blog *model.Blog) (*model.Blog, error) {
	result := configuration.DB.Updates(blog)
	if err := result.Error; err != nil {
		log.Printf("Update Blog failed, blog: %v, error: %v\n", *blog, err)
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) DeleteBlog(ctx context.Context, id int) (*model.Blog, error) {
	var blog *model.Blog
	if err := configuration.DB.Delete(&blog, id).Error; err != nil {
		log.Printf("Delete Blog failed, id: %v, blog: %v, error: %v\n", id, *blog, err)
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) GetBlogById(ctx context.Context, id int) (*model.Blog, error) {
	var blog model.Blog
	if err := configuration.DB.Where("id = ?", id).Find(&blog).Error; err != nil {
		log.Printf("Get Blog failed, id: %v, blog: %v, error: %v\n", id, blog, err)
		return nil, err
	}
	return &blog, nil
}

func (r *BlogRepository) ListBlogWithPagination(ctx context.Context, pageNum int) ([]model.Blog, error) {
	var listBlog []model.Blog
	if err := configuration.DB.Limit(pageNum).Order("ctime desc").Find(&listBlog).Error; err != nil {
		log.Printf("List Blog failed, pageNum: %v, blogList: %v, error: %v", pageNum, listBlog, err)
		return nil, err
	}
	return listBlog, nil
}

func (r *BlogRepository) CreateBlogLog(ctx context.Context, blogLog *model.BlogLog) error {
	result := configuration.DB.Create(blogLog)
	if err := result.Error; err != nil {
		log.Printf("Create Blog Log failed, blogLog: %v, error: %v\n", *blogLog, err)
		return err
	}
	if rowNum := result.RowsAffected; rowNum == 0 {
		err := errors.New("create Blog Log failed, rowNum: 0")
		log.Printf("Create Blog Log Failed, blogLog: %v, error: %v", *blogLog, err)
		return err
	}
	return nil
}

func (r *BlogRepository) GetBlogLog(ctx context.Context, id int) (*model.BlogLog, error) {
	var blogLog model.BlogLog
	if err := configuration.DB.Where("id = ?", id).Find(&blogLog).Error; err != nil {
		log.Printf("Get Blog Log failed, blogLog: %v, id: %v error: %v\n", blogLog, id, err)
		return nil, err
	}
	return &blogLog, nil
}
