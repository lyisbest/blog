package controller

import (
	"blog/apps/blog/request"
	"blog/apps/blog/service"
	"blog/apps/user/constant"
	"blog/constants"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
)

type BlogController struct {
	blogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

func (c *BlogController) CreateBlog(ctx *gin.Context, blogCreateRequest request.BlogCreateRequest) (interface{}, error) {
	user, exists := ctx.Get("username")
	if !exists {
		return nil, constant.UserNotExistError
	}
	userName, ok := user.(string)
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

func (c *BlogController) DeleteBlog(ctx *gin.Context, blogDeleteRequest request.BlogDeleteRequest) (interface{}, error) {
	user, exists := ctx.Get("username")
	if !exists {
		return nil, constant.UserNotExistError
	}
	userName, ok := user.(string)
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

func (c *BlogController) UpdateBlog(ctx *gin.Context, blogUpdateRequest request.BlogUpdateRequest) (interface{}, error) {
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

func (c *BlogController) GetBlog(ctx *gin.Context, blogGetRequest request.BlogGetRequest) (interface{}, error) {
	blog, err := c.blogService.GetBlog(ctx, blogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog failed, blog: %v, error: %v", *blog, err)
		return nil, err
	}
	return *blog, nil
}

func (c *BlogController) ListBlogWithPagination(ctx *gin.Context, listBlogRequest request.ListBlogRequest) (interface{}, error) {
	listBlog, err := c.blogService.ListBlogWithPagination(ctx, listBlogRequest.PageNum, listBlogRequest.PageSize)
	if err != nil {
		log.Printf("List Blog failed, listBlog: %v, error: %v", listBlog, err)
		return nil, err
	}
	return listBlog, nil
}

func (c *BlogController) GetBlogLog(ctx *gin.Context, blogLogGetRequest request.BlogLogGetRequest) (interface{}, error) {
	blogLog, err := c.blogService.GetBlogLog(ctx, blogLogGetRequest.Id)
	if err != nil {
		log.Printf("Get Blog Log failed, blog: %v, error: %v", *blogLog, err)
		return nil, err
	}
	return *blogLog, nil
}

func (c *BlogController) AddBlogView(ctx *gin.Context, blogGetRequest request.BlogGetRequest) (interface{}, error) {
	return nil, nil
}

func (c *BlogController) UpdateImage(ctx *gin.Context) {
	image, err := ctx.FormFile("file")
	if err != nil {
		utils.EndWithError(ctx, err)
		return
	}
	extName := path.Ext(image.Filename)
	if extName != ".bmp" && extName != ".jpg" && extName != ".jpeg" && extName != ".png" && extName != ".gif" && extName != ".webp" {
		utils.EndWithError(ctx, err)
		return
	}
	fileName := "image_" + utils.GetUUID() + extName
	//test in develop
	realPath := "C:/Users/11873/Desktop/images/"
	//live: in linux server
	//realPath := "/home/ubuntu/images/"
	nginxPath := "/images/"
	pathExist, err := utils.PathExists(realPath)
	if err != nil {
		log.Printf("%v", err)
	}
	if !pathExist {
		err = os.MkdirAll(realPath, os.ModePerm)
		if err != nil {
			log.Printf("%v", err)
		}
	}

	err = ctx.SaveUploadedFile(image, realPath+fileName)
	if err != nil {
		utils.EndWithError(ctx, err)
		return
	}
	//utils.ResponseWithData(ctx, "http://"+utils.GetOutboundIP().String()+nginxPath+fileName)
	utils.ResponseWithData(ctx, "http://43.142.250.217"+nginxPath+fileName)
}
