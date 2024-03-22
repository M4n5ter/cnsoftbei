# shell
set windows-shell := ["powershell.exe", "-c"]

#====================================== alias start ============================================#

alias b := build
alias r := run
alias t := test
alias deps := dependencies

#======================================= alias end =============================================#

#===================================== targets start ===========================================#

# default target
default: lint test

# PLEASE DO THIS FIRSET!
[unix]
add-hook:
    @echo "just" > {{pre_commit}}
    @chmod +x {{pre_commit}}

[windows]
add-hook:
    @echo "just" > {{pre_commit}}

# generate router, api, service code.
[confirm("""
Are you sure you want to generate router, api, service code?
你确定要生成 router, api, service 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
logic targets:
    @cd {{join(root,"cmd")}} {{and}} go run . logic {{targets}} -d {{root}}

# generate router code.
[confirm("""
Are you sure you want to generate router code?
你确定要生成 router 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
router targets:
    @cd {{join(root,"cmd")}} {{and}} go run . router {{targets}} -d {{join(root,"core","router")}}

# generate api code.
[confirm("""
Are you sure you want to generate api code?
你确定要生成 api 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
api targets:
    @cd {{join(root,"cmd")}} {{and}} go run . api {{targets}} -d {{join(root,"core","api")}}

# generate service code.
[confirm("""
Are you sure you want to generate service code?
你确定要生成 service 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
service targets:
    @cd {{join(root,"cmd")}} {{and}} go run . service {{targets}} -d {{join(root,"core","service")}}

# generate middleware code.
[confirm("""
Are you sure you want to generate middleware code?
你确定要生成 middleware 代码吗？
input 'Y/N' to continue or exit.
输入 'Y/N' 继续或退出。
""")]
middleware targets:
    @cd {{join(root,"cmd")}} {{and}} go run . middleware {{targets}} -d {{join(root,"core","middleware")}}

# go build
build: swag
    @echo "Building..."
    @GIN_MODE=release go build -tags=jsoniter -ldflags "-s -w" -o {{bin}} {{main_file}}
    @echo "Build done."

# go run
run: swag
    @go run -ldflags "-X 'main.logLevel=debug'" -tags=swagger {{main_file}}

# go test
test:
    @go test -v {{join(".", "...")}}

# generate swagger docs
swag: dep-swag
    @cd {{server}} {{and}} swag init -g swagger.go

# lint
lint: dep-golangci-lint
    @go mod tidy 
    @golangci-lint run

# run openobserve
oo:
    ZO_ROOT_USER_EMAIL="root@example.com" ZO_ROOT_USER_PASSWORD="Complexpass#123" {{join(root, "openobserve")}}

# run redis/dragonfly
redis:
    {{join(root, "dragonfly")}} --dir {{join(root, "data", "redis")}} --logtostderr --requirepass=youshallnotpass --cache_mode=true -dbnum 1 --bind 0.0.0.0 --port 6379  --snapshot_cron "*/30 * * * *" --maxmemory=12gb --keys_output_limit=12288

# install dependencies
dependencies: dep-swag dep-golangci-lint dep-gofumpt

# a tool to help you write API docs
dep-swag:
    @go install github.com/swaggo/swag/cmd/swag@latest

# a linter for Go
dep-golangci-lint:
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# a stricter gofmt
dep-gofumpt:
    @go install mvdan.cc/gofumpt@latest

#===================================== targets end ===========================================#

#=================================== variables start =========================================#
# project name
project_name := "cnsoftbei"

# project root directory
root := justfile_directory()

# binary path
bin := join(root, project_name)

# main.go path
main_file := join(root, "server", "main.go")

# server path
server := join(root, "server")

# pre-commit path
pre_commit := join(root, ".git", "hooks", "pre-commit")

#=================================== variables end =========================================#

and := if os_family() == "windows" {";"} else {"&&"}