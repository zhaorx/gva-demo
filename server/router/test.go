package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("test").Use(middleware.OperationRecord())
	{
		UserRouter.POST("testT", v1.TestT)
	}
}
