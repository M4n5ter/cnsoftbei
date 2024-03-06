package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m4n5ter/cnsoftbei/core/service"
)

type UserAPI struct{}

func (ua *UserAPI) List(ctx *gin.Context) {
	userService := service.Group.UserService
	_ = userService.List()
}
