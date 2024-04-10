# set shell
set windows-shell := ["powershell.exe", "-c"]

# set `&&` or `;` for different OS
and := if os_family() == "windows" {";"} else {"&&"}

# load environment from `.env` file
set dotenv-load

#====================================== alias start ============================================#

alias b := build
alias r := run
alias t := test
alias deps := dependencies
alias new := new-migration
alias up := up-migration
alias up1 := up-by-one-migration
alias down := down-migration
alias down1 := down-by-one-migration
alias version := version-migration
alias status := status-migration

#======================================= alias end =============================================#

#===================================== targets start ===========================================#

# default target - `just` 默认目标
default: lint test

# PLEASE DO THIS FIRSET! - 务必先执行 `just add-hook`
[unix]
add-hook:
    @echo "just" > {{pre_commit}}
    @chmod +x {{pre_commit}}

[windows]
add-hook:
    @echo "just" > {{pre_commit}}

# generate router, api, service code. - 生成 router, api, service 代码
[confirm("""
Are you sure you want to generate router, api, service code?
你确定要生成 router, api, service 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
logic targets:
    @cd {{codegen}} {{and}} go run . logic {{targets}} -d {{root}}

# generate router code. - 生成 router 代码
[confirm("""
Are you sure you want to generate router code?
你确定要生成 router 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
router targets:
    @cd {{codegen}} {{and}} go run . router {{targets}} -d {{join(root,"core","router")}}

# generate api code. - 生成 api 代码
[confirm("""
Are you sure you want to generate api code?
你确定要生成 api 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
api targets:
    @cd {{codegen}} {{and}} go run . api {{targets}} -d {{join(root,"core","api")}}

# generate service code. - 生成 service 代码
[confirm("""
Are you sure you want to generate service code?
你确定要生成 service 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
service targets:
    @cd {{codegen}} {{and}} go run . service {{targets}} -d {{join(root,"core","service")}}

# generate middleware code. - 生成 middleware 代码
[confirm("""
Are you sure you want to generate middleware code?
你确定要生成 middleware 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
middleware targets:
    @cd {{codegen}} {{and}} go run . middleware {{targets}} -d {{join(root,"core","middleware")}}

# go build
[unix]
build:
    @echo "Building..."
    @GIN_MODE=release go build -tags="jsoniter no_swagger" -ldflags "-s -w" -o {{bin}} {{main_file}}
    @echo "Build done."

[windows]
build:
    @echo "Building..."
    @$env:GIN_MODE="release" {{and}} go build -tags="jsoniter no_swagger" -ldflags "-s -w" -o {{bin}} {{main_file}}
    @echo "Build done."

# go run
run: swag
    @go run -ldflags "-X 'main.logLevel=debug'" {{main_file}}

# go test
test:
    @go test -v {{join(".", "...")}}

# generate swagger docs - 生成 swagger 文档
swag: dep-swag
    @cd {{server}} {{and}} swag init -g swagger.go

# lint - 代码检查
lint: dep-golangci-lint
    @go mod tidy 
    @golangci-lint run

# new a migration file with specified name - 新建一个指定名称的迁移文件
new-migration name: dep-goose
    @goose create {{name}} sql

# run migration - 运行所有的迁移文件
up-migration: dep-goose
    @goose up

# run migration by one - 运行一个迁移文件
up-by-one-migration: dep-goose
    @goose up-by-one

# rollback all migrations - 回滚所有的迁移文件
down-migration: dep-goose
    @goose reset

# rollback one migration - 回滚一个迁移文件
down-by-one-migration: dep-goose
    @goose down

# show current migration version - 显示当前迁移版本
version-migration: dep-goose
    @goose version

# show current migration status - 显示当前迁移状态
status-migration: dep-goose
    @goose status

# install dependencies - 安装依赖工具
dependencies: dep-goose dep-swag dep-golangci-lint dep-gofumpt

# a lightweight, framework-independent database migration tool - 一个轻量级、独立于框架的数据库迁移工具
dep-goose:
    @go install github.com/pressly/goose/v3/cmd/goose@latest

# a tool to help you write API docs - 一个帮助你编写 API 文档的工具
dep-swag:
    @go install github.com/swaggo/swag/cmd/swag@latest

# a linter for Go - 一个 Go 语言的代码检查工具
dep-golangci-lint:
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# a stricter gofmt - 一个更严格的 gofmt
dep-gofumpt:
    @go install mvdan.cc/gofumpt@latest

#===================================== targets end ===========================================#

#=================================== variables start =========================================#

# project name - 项目名称
project_name := "cnsoftbei"

# project root directory - 项目根目录
root := justfile_directory()

# binary path - go build 输出的二进制文件路径
bin := join(root, project_name)

# main.go path - main.go 文件路径
main_file := join(root, "cmd", "server", "main.go")

# server path - server 目录路径
server := join(root, "cmd", "server")

# codegen path - codegen 目录路径
codegen := join(root, "cmd", "codegen")

# pre-commit path - pre-commit 文件路径
pre_commit := join(root, ".git", "hooks", "pre-commit")

#=================================== variables end =========================================#