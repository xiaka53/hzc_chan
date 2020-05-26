package router

import (
	_ "api/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//路由初始化
func InitRouter(middlewares ...gin.HandlerFunc) (router *gin.Engine) {
	router = gin.Default()
	router.Use(middlewares...)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}
