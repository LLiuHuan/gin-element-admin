package router

import "github.com/gin-gonic/gin"

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{

	}
	return BaseRouter
}
