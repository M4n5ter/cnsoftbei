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

# go mod tidy
tidy target:
    @go mod tidy

# generate swagger docs
swag: dep-swag
    @cd {{server}} {{and}} swag init -g swagger.go

# lint
lint: dep-golangci-lint
    @golangci-lint run

# run openobserve
oo:
    ZO_ROOT_USER_EMAIL="root@example.com" ZO_ROOT_USER_PASSWORD="Complexpass#123" {{join(root, "openobserve")}}

redis:
    {{join(root, "dragonfly")}} --dir {{join(root, "data", "redis")}} --logtostderr --requirepass=youshallnotpass --cache_mode=true -dbnum 1 --bind localhost --port 6379  --snapshot_cron "*/30 * * * *" --maxmemory=12gb --keys_output_limit=12288

# install dependencies
dependencies: dep-swag dep-golangci-lint dep-gofumpt

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

#=================================== variables end =========================================#

and := if os_family() == "windows" {";"} else {"{{and}}"}