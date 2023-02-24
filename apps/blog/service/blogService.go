package service

import (
	"blog/apps/blog/constant"
	"blog/apps/blog/model"
	"blog/apps/blog/repository"
	"blog/configuration"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type BlogService struct {
	blogRepository repository.BlogRepository
}

func NewBlogService(blogRepository repository.BlogRepository) *BlogService {
	return &BlogService{blogRepository: blogRepository}
}

func (s *BlogService) CreateBlog(ctx *gin.Context, title, content, userName string) error {
	now := time.Now().Unix()
	blog := &model.Blog{
		Title:   title,
		Content: content,
		Creator: userName,
		Ctime:   now,
		Mtime:   now,
	}
	blogLog := &model.BlogLog{
		Operator:      userName,
		OperationType: constant.Create,
		Operation:     "",
		Ctime:         now,
	}
	err := configuration.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		blog, err := s.blogRepository.CreateBlog(ctx, blog)
		if err != nil {
			return err
		}

		blogLog.BlogId = blog.Id
		if err := s.blogRepository.CreateBlogLog(ctx, blogLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) UpdateBlog(ctx *gin.Context, title, content, userName string, id int) error {
	now := time.Now().Unix()
	blog := &model.Blog{
		Id:      id,
		Title:   title,
		Content: content,
		Mtime:   now,
	}
	blogLog := &model.BlogLog{
		BlogId:        id,
		Operator:      userName,
		OperationType: constant.Update,
		Operation:     "",
		Ctime:         now,
	}

	err := configuration.DB.Transaction(func(tx *gorm.DB) error {
		if _, err := s.blogRepository.UpdateBlog(ctx, blog); err != nil {
			return err
		}

		if err := s.blogRepository.CreateBlogLog(ctx, blogLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) DeleteBlog(ctx *gin.Context, id int, userName string) error {
	now := time.Now().Unix()
	blogLog := &model.BlogLog{
		BlogId:        id,
		Operator:      userName,
		OperationType: constant.Delete,
		Operation:     "",
		Ctime:         now,
	}

	err := configuration.DB.Transaction(func(tx *gorm.DB) error {
		if _, err := s.blogRepository.DeleteBlog(ctx, id); err != nil {
			return err
		}

		if err := s.blogRepository.CreateBlogLog(ctx, blogLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) GetBlog(ctx *gin.Context, id int) (*model.Blog, error) {
	blog, err := s.blogRepository.GetBlogById(ctx, id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *BlogService) ListBlogWithPagination(ctx *gin.Context, pageNum int) ([]model.Blog, error) {
	return s.blogRepository.ListBlogWithPagination(ctx, pageNum)
}

func (s *BlogService) GetBlogLog(ctx *gin.Context, id int) (*model.BlogLog, error) {
	blogLog, err := s.blogRepository.GetBlogLog(ctx, id)
	if err != nil {
		return nil, err
	}
	return blogLog, nil
}
