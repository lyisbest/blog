package controller

import (
	"github.com/gin-gonic/gin"
)

type SysController struct {
}

func NewSysController() *SysController {
	return &SysController{}
}

// Ping healthy check
func (c *SysController) Ping(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}
