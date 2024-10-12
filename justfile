set shell := ["powershell.exe", "-c"]

default:
  @just --justfile {{justfile()}} --list

# Run acceptance tests
test:
	$env:TF_ACC=1; go test ./... -v -timeout 120m

# Generate documentation
generate:
	go generate ./...

# Run and fix static checks
lint:
	golangci-lint run --fix
