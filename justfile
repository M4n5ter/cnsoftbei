# shell
set windows-shell := ["cmd.exe", "/c"]

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
    @go run -tags=swagger {{main_file}}

# go test
test:
    @go test -v {{join(server, "...")}}

# go mod tidy
tidy target:
    @echo "Tidying {{target}}..."
    @cd {{target}} && go mod tidy

# generate swagger docs
swag: dep-swag
    @cd {{server}} && swag init -g swagger.go

# lint
lint: dep-golangci-lint
    golangci-lint run {{lint_dirs}}

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

# golangci-lint dirs
lint_dirs := join(server, "...") + " " + join(root, "core", "...") + " " + join(root, "common", "...")

#=================================== variables end =========================================#
