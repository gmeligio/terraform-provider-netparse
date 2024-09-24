default: testacc

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

generate:
	go generate ./...

# Run acceptance tests
.PHONY: generate testacc
