package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/m4n5ter/cnsoftbei/cmd/server/docs"
	"github.com/m4n5ter/cnsoftbei/core/middleware"
	"github.com/m4n5ter/cnsoftbei/core/router"
	"github.com/m4n5ter/cnsoftbei/pkg/config"
	"github.com/m4n5ter/cnsoftbei/pkg/yalog"
)

var (
	configPath string

	// this will be set by the linker
	logLevel string
)

func main() {
	flag.StringVar(&configPath, "c", "config.toml", "path to the configuration file")
	flag.Parse()

	switch logLevel {
	case "debug":
		yalog.SetLevelDebug()
	case "info":
		yalog.SetLevelInfo()
	case "warn":
		yalog.SetLevelWarn()
	case "error":
		yalog.SetLevelError()
	default:
		yalog.SetLevelInfo()
	}

	// load configuration from file
	conf := config.MustLoad(configPath)

	r := gin.Default()
	if err := r.SetTrustedProxies(conf.TrustedProxies); err != nil {
		yalog.Fatalf("failed to set trusted proxies: %v", err)
	}

	r.Use(middleware.Redis(conf))

	// Register the router
	router.Routes.Register(&r.RouterGroup)

	if err := r.Run(fmt.Sprintf("%s:%v", conf.Host, conf.Port)); err != nil {
		yalog.Fatalf("failed to start server: %v", err)
	}
}
