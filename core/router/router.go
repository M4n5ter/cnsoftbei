package router

import "github.com/gin-gonic/gin"

// Routes contains all routers
var Routes = &routes{routers: make([]router, 0)}

type routes struct {
	routers []router
}

// register all routers
func (r *routes) Register(rg *gin.RouterGroup) {
	for _, router := range r.routers {
		router.Register(rg)
	}
}

type router interface {
	Register(*gin.RouterGroup)
}
