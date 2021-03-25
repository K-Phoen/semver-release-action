IMAGE=krogon/semver-release-action
VERSION?=latest

.PHONY: lint
lint: vendor
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:latest golangci-lint run -c .golangci.yaml --skip-dirs-use-default

.PHONY: test
test: vendor
	go test -mod=vendor ./...

vendor:
	go mod vendor

.PHONY: build
build:
	go build -o bumper

.PHONY: image
image:
	docker build -f Dockerfile -t $(IMAGE):$(VERSION) .

.PHONY: push
push: image
	docker push $(IMAGE):$(VERSION)

.PHONY: update_action
update_action:
	sed "s/\*\*VERSION_PLACEHOLDER\*\*/$(VERSION)/" action.yml.dist > action.yml

.PHONY: clean
clean:
	rm bumper
	rm -rf vendor
