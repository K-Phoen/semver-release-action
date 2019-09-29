.PHONY: lint
lint: vendor
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:latest golangci-lint run -c .golangci.yaml --skip-dirs-use-default

.PHONY: test
test: vendor
	go test -mod=vendor ./...

vendor:
	go mod vendor