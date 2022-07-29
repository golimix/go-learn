package handler

import (
	"github.com/golimix/go-learn/framework/gin/casbin/component"
	"github.com/gin-gonic/gin"
)

// handler/resource_handler.go

func ReadResource(c *gin.Context) {
	c.JSON(200, component.RestResponse{Code: 1, Message: "read resource successfully", Data: "resource"})
}

func WriteResource(c *gin.Context) {
	c.JSON(200, component.RestResponse{Code: 1, Message: "write resource successfully", Data: "resource"})
}
