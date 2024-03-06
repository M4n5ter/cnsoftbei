package router

import "github.com/gin-gonic/gin"

type router interface {
	Register(*gin.RouterGroup)
}

type routes struct {
	routers []router
}

// register all routers
func (r *routes) Register(rg *gin.RouterGroup) {
	for _, router := range r.routers {
		router.Register(rg)
	}
}

func init() {
	Routes.routers = make([]router, 0)
}

// Routes contains all routers
var Routes = new(routes)
