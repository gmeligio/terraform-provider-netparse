default: testacc

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

generate:
	go generate ./...

lint:
	golangci-lint run --fix

# Run acceptance tests
.PHONY: generate lint testacc
