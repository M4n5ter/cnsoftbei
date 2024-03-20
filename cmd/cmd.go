package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/m4n5ter/cnsoftbei/cmd/core/api"
	"github.com/m4n5ter/cnsoftbei/cmd/core/middleware"
	"github.com/m4n5ter/cnsoftbei/cmd/core/router"
	"github.com/m4n5ter/cnsoftbei/cmd/core/service"
	"github.com/spf13/cobra"
)

func init() {
	logicCmd.Flags().StringVarP(&dir, "dir", "d", "", "the dir to place the generated files (required)")
	logicCmd.MarkFlagRequired("dir")

	routerCmd.Flags().StringVarP(&dir, "dir", "d", "", "the dir to place the generated files (required)")
	routerCmd.MarkFlagRequired("dir")

	serviceCmd.Flags().StringVarP(&dir, "dir", "d", "", "the dir to place the generated files (required)")
	serviceCmd.MarkFlagRequired("dir")

	apiCmd.Flags().StringVarP(&dir, "dir", "d", "", "the dir to place the generated files (required)")
	apiCmd.MarkFlagRequired("dir")

	middlewareCmd.Flags().StringVarP(&dir, "dir", "d", "", "the dir to place the generated files (required)")
	middlewareCmd.MarkFlagRequired("dir")

	rootCmd.AddCommand(logicCmd)
	rootCmd.AddCommand(routerCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(middlewareCmd)
}

var dir string

var (
	rootCmd = &cobra.Command{
		Use:   "gen",
		Short: "Code generator",
	}

	logicCmd = &cobra.Command{
		Use:   "logic ...",
		Short: "Generate router, service, api code.",
		Run: func(cmd *cobra.Command, args []string) {
			logicRun(args, dir)
		},
	}

	routerCmd = &cobra.Command{
		Use:   "router ...",
		Short: "Generate router code",
		Run: func(cmd *cobra.Command, args []string) {
			routerRun(args, dir)
		},
	}

	serviceCmd = &cobra.Command{
		Use:   "service ...",
		Short: "Generate service code",
		Run: func(cmd *cobra.Command, args []string) {
			serviceRun(args, dir)
		},
	}

	apiCmd = &cobra.Command{
		Use:   "api ...",
		Short: "Generate api code",
		Run: func(cmd *cobra.Command, args []string) {
			apiRun(args, dir)
		},
	}

	middlewareCmd = &cobra.Command{
		Use:   "middleware ...",
		Short: "Generate middleware code",
		Run: func(cmd *cobra.Command, args []string) {
			middlewareRun(args, dir)
		},
	}
)

func logicRun(names []string, dir string) {
	if len(names) < 1 {
		fmt.Println("Please provide a name for the logic")
		return
	}

	routerDir := filepath.Join(dir, "core", "router")
	serviceDir := filepath.Join(dir, "core", "service")
	apiDir := filepath.Join(dir, "core", "api")
	os.MkdirAll(routerDir, 0o755)
	os.MkdirAll(serviceDir, 0o755)
	os.MkdirAll(apiDir, 0o755)

	routerRun(names, routerDir)
	serviceRun(names, serviceDir)
	apiRun(names, apiDir)
}

func routerRun(names []string, dir string) {
	if len(names) < 1 {
		fmt.Println("Please provide a name for the router")
		return
	}

	os.MkdirAll(dir, 0o755)

	for _, name := range names {
		file, err := os.Create(filepath.Join(dir, name+"_router.go"))
		if err != nil {
			panic(err)
		}
		err = router.New(f, name).Generate(file)
		if err != nil {
			panic(err)
		}
	}
}

func serviceRun(names []string, dir string) {
	if len(names) < 1 {
		fmt.Println("Please provide a name for the service")
		return
	}

	os.MkdirAll(dir, 0o755)

	for _, name := range names {
		file, err := os.Create(filepath.Join(dir, name+"_service.go"))
		if err != nil {
			panic(err)
		}
		err = service.New(f, name).Generate(file)
		if err != nil {
			panic(err)
		}
	}
}

func apiRun(names []string, dir string) {
	if len(names) < 1 {
		fmt.Println("Please provide a name for the api")
		return
	}

	os.MkdirAll(dir, 0o755)

	for _, name := range names {
		file, err := os.Create(filepath.Join(dir, name+"_api.go"))
		if err != nil {
			panic(err)
		}
		err = api.New(f, name).Generate(file)
		if err != nil {
			panic(err)
		}
	}
}

func middlewareRun(names []string, dir string) {
	if len(names) < 1 {
		fmt.Println("Please provide a name for the middleware")
		return
	}

	os.MkdirAll(dir, 0o755)

	for _, name := range names {
		file, err := os.Create(filepath.Join(dir, name+".go"))
		if err != nil {
			panic(err)
		}
		err = middleware.New(f, name).Generate(file)
		if err != nil {
			panic(err)
		}
	}
}
