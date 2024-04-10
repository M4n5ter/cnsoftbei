//go:build !no_pprof

package router

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func init() {
	Routes.routers = append(Routes.routers, &pprofRouter{})
}

type pprofRouter struct{}

func (pr *pprofRouter) Register(rg *gin.RouterGroup) {
	pg := rg.Group("/debug/pprof")
	{
		pg.GET("/", gin.WrapF(pprof.Index))
		pg.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		pg.GET("/profile", gin.WrapF(pprof.Profile))
		pg.POST("/symbol", gin.WrapF(pprof.Symbol))
		pg.GET("/symbol", gin.WrapF(pprof.Symbol))
		pg.GET("/trace", gin.WrapF(pprof.Trace))
		pg.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		pg.GET("/block", gin.WrapH(pprof.Handler("block")))
		pg.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		pg.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		pg.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		pg.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
