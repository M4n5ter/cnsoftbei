package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m4n5ter/cnsoftbei/common/log"
	"github.com/m4n5ter/cnsoftbei/core/router"
	"github.com/m4n5ter/cnsoftbei/server/config"
	_ "github.com/m4n5ter/cnsoftbei/server/docs"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "c", "config.toml", "path to the configuration file")
	flag.Parse()

	// load configuration from file
	conf := config.MustLoad(configPath)

	r := gin.Default()
	if err := r.SetTrustedProxies(conf.TrustedProxies); err != nil {
		log.Panicf("failed to set trusted proxies: %v", err)
	}

	// Register the router
	router.Routes.Register(&r.RouterGroup)

	if err := r.Run(fmt.Sprintf("%s:%v", conf.Host, conf.Port)); err != nil {
		log.Panicf("failed to start server: %v", err)
	}
}
