package controller

import (
	"blog/apps/blog/request"
	"blog/apps/blog/service"
	"blog/apps/user/constant"
	"blog/constants"
	"github.com/gin-gonic/gin"
	"log"
)

type BlogController struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

func (c *BlogController) CreateBlog(ctx *gin.Context) (interface{}, error) {
	var blogCreateRequest request.BlogCreateRequest
	if err := ctx.ShouldBindJSON(&blogCreateRequest); err != nil {
		return nil, constants.ResolveError
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		return nil, constant.UserNotExistError
	}
	userName, ok := cookie.(string)
	if !ok {
		return nil, constants.ResolveError
	}
	err := c.blogService.CreateBlog(ctx, blogCreateRequest.Title, blogCreateRequest.Content, userName)
	if err != nil {
		log.Printf("Create Blog failed, error: %v", err)
		return nil, err
	}
	return nil, nil
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) (interface{}, error) {
	var blogDeleteRequest request.BlogDeleteRequest
	if err := ctx.ShouldBindJSON(&blogDeleteRequest); err != nil {
		return nil, constants.ResolveError
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		return nil, constant.UserNotExistError
	}
	userName, ok := cookie.(string)
	if !ok {
		return nil, constants.ResolveError
	}
	err := c.blogService.DeleteBlog(ctx, blogDeleteRequest.Id, userName)
	if err != nil {
		log.Printf("Delete Blog failed, error: %v", err)
		return nil, err
	}
	return nil, nil
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) (interface{}, error) {
	var blogUpdateRequest request.BlogUpdateRequest
	if err := ctx.ShouldBindJSON(&blogUpdateRequest); err != nil {
		return nil, constants.ResolveError
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		return nil, constant.UserNotExistError
	}
	userName, ok := cookie.(string)
	if !ok {
		return nil, constants.ResolveError
	}
	err := c.blogService.UpdateBlog(ctx, blogUpdateRequest.Title, blogUpdateRequest.Content, userName, blogUpdateRequest.Id)
	if err != nil {
		log.Printf("Update Blog failed, error: %v", err)
		return nil, err
	}
	return nil, nil
}

func (c *BlogController) GetBlog(ctx *gin.Context) (interface{}, error) {
	var blogGetRequest request.BlogGetRequest
	if err := ctx.ShouldBindJSON(&blogGetRequest); err != nil {
		return nil, constants.ResolveError
	}
	blog, err := c.blogService.GetBlog(ctx, blogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog failed, blog: %v, error: %v", *blog, err)
		return nil, err
	}
	return *blog, nil
}

func (c *BlogController) ListBlogWithPagination(ctx *gin.Context) (interface{}, error) {
	var listBlogRequest request.ListBlogRequest
	if err := ctx.ShouldBindJSON(&listBlogRequest); err != nil {
		return nil, constants.ResolveError
	}
	listBlog, err := c.blogService.ListBlogWithPagination(ctx, listBlogRequest.PageNum)
	if err != nil {
		log.Printf("List Blog failed, listBlog: %v, error: %v", listBlog, err)
		return nil, err
	}
	return listBlog, nil
}

func (c *BlogController) GetBlogLog(ctx *gin.Context) (interface{}, error) {
	var blogLogGetRequest request.BlogLogGetRequest
	if err := ctx.ShouldBindJSON(&blogLogGetRequest); err != nil {
		return nil, constants.ResolveError
	}
	blogLog, err := c.blogService.GetBlogLog(ctx, blogLogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog Log failed, blog: %v, error: %v", *blogLog, err)
		return nil, err
	}
	return *blogLog, nil
}
