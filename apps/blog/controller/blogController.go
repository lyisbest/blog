package controller

import (
	constant2 "blog/apps/blog/constant"
	"blog/apps/blog/request"
	"blog/apps/blog/service"
	"blog/apps/user/constant"
	"blog/constants"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type BlogController struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blogCreateRequest request.BlogCreateRequest
	if err := ctx.ShouldBindJSON(&blogCreateRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		utils.ResponseWithError(ctx, constant.UserNotExistError)
		return
	}
	userName, ok := cookie.(string)
	if !ok {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	err := c.blogService.CreateBlog(ctx, blogCreateRequest.Title, blogCreateRequest.Content, userName)
	if err != nil {
		log.Printf("Create Blog failed, error: %v", err)
		utils.ResponseWithError(ctx, constant2.BlogCreateFailError)
		return
	}
	utils.ResponseNormal(ctx)
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	var blogDeleteRequest request.BlogDeleteRequest
	if err := ctx.ShouldBindJSON(&blogDeleteRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		utils.ResponseWithError(ctx, constant.UserNotExistError)
		return
	}
	userName, ok := cookie.(string)
	if !ok {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	err := c.blogService.DeleteBlog(ctx, blogDeleteRequest.Id, userName)
	if err != nil {
		log.Printf("Delete Blog failed, error: %v", err)
		utils.ResponseWithError(ctx, constant2.BlogDeleteFailError)
		return
	}
	utils.ResponseNormal(ctx)
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	var blogUpdateRequest request.BlogUpdateRequest
	if err := ctx.ShouldBindJSON(&blogUpdateRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	cookie, exists := ctx.Get("username")
	if !exists {
		utils.ResponseWithError(ctx, constant.UserNotExistError)
		return
	}
	userName, ok := cookie.(string)
	if !ok {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	err := c.blogService.UpdateBlog(ctx, blogUpdateRequest.Title, blogUpdateRequest.Content, userName, blogUpdateRequest.Id)
	if err != nil {
		log.Printf("Update Blog failed, error: %v", err)
		utils.ResponseWithError(ctx, constant2.BlogUpdateFailError)
		return
	}
	utils.ResponseNormal(ctx)
}

func (c *BlogController) GetBlog(ctx *gin.Context) {
	var blogGetRequest request.BlogGetRequest
	if err := ctx.ShouldBindJSON(&blogGetRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	blog, err := c.blogService.GetBlog(ctx, blogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog failed, blog: %v, error: %v", *blog, err)
		utils.ResponseWithError(ctx, constant2.BlogGetFailError)
		return
	}
	utils.ResponseWithData(ctx, *blog)
}

func (c *BlogController) ListBlogWithPagination(ctx *gin.Context) {
	var listBlogRequest request.ListBlogRequest
	if err := ctx.ShouldBindJSON(&listBlogRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	listBlog, err := c.blogService.ListBlogWithPagination(ctx, listBlogRequest.PageNum)
	if err != nil {
		log.Printf("List Blog failed, listBlog: %v, error: %v", listBlog, err)
		utils.ResponseWithError(ctx, constant2.ListBlogGetFailError)
	}
	utils.ResponseWithData(ctx, listBlog)
}

func (c *BlogController) GetBlogLog(ctx *gin.Context) {
	var blogLogGetRequest request.BlogLogGetRequest
	if err := ctx.ShouldBindJSON(&blogLogGetRequest); err != nil {
		utils.ResponseWithError(ctx, constants.ResolveError)
		return
	}
	blogLog, err := c.blogService.GetBlogLog(ctx, blogLogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog Log failed, blog: %v, error: %v", *blogLog, err)
		utils.ResponseWithError(ctx, constant2.BlogLogGetFailError)
		return
	}
	utils.ResponseWithData(ctx, *blogLog)
}
