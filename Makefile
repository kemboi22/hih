.PHONY: setup-hooks fmt vet test

setup-hooks:
	git config core.hooksPath .githooks
	@echo "Git hooks configured — .githooks/ is now active."

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test -v ./...
