package main

import (
	"fmt"
	"log"

	"github.com/24king/go-learn/framework/gin/casbin/component"
	"github.com/24king/go-learn/framework/gin/casbin/handler"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(component.DB)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Initialize Gin router
	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig)) // CORS configuraion
	router.POST("/user/login", handler.Login)
	// Secure our API
	resource := router.Group("/api")
	{
		resource.GET("/resource", middleware.Authorize("resource", "read", adapter), handler.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", adapter), handler.WriteResource)
	}
}

func main() {
	defer component.DB.Close()

	// Start our application
	err := router.Run(":8081")
	if err != nil {
		panic(fmt.Sprintf("failed to start gin engin: %v", err))
	}
	log.Println("application is now running...")
}
