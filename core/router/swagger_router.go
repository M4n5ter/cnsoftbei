//go:build !no_swagger

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	Routes.routers = append(Routes.routers, &swaggerRouter{})
}

type swaggerRouter struct{}

func (sr *swaggerRouter) Register(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
