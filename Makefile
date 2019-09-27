.PHONY: lint
lint: vendor
	docker run --rm -v $(CURDIR):/app -w /app/bumper golangci/golangci-lint:latest golangci-lint run --skip-dirs-use-default

.PHONY: test
test: vendor
	cd bumper && go test -mod=vendor ./...

vendor:
	cd bumper && go mod vendor