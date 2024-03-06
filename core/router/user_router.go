package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m4n5ter/cnsoftbei/core/api"
)

func init() {
	Routes.routers = append(Routes.routers, &userRouter{})
}

type userRouter struct{}

func (ur *userRouter) Register(rg *gin.RouterGroup) {
	userAPI := api.Group.UserAPI

	ug := rg.Group("/users")
	{
		ug.GET("", userAPI.List)
	}
}
