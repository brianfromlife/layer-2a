
GO := GO111MODULE=on go
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOMIG ?= $(GOBIN)/migrate
GOLINT ?= $(GOBIN)/golint
GOSTATIC ?= $(GOBIN)/staticcheck

deps:
	go mod tidy
	go mod vendor

down:
	@$(GOMIG) -verbose -source file://migration -database sqlite3://sqlite3/layer_2.db force 1 down all || (echo "error exporting database"; exit 1)

up:
	$(GOMIG) -verbose -source file://migration -database sqlite3://sqlite3/layer_2.db up 1 version

	
lint:
	@$(GOSTATIC) ./... | grep -v vendor/ && exit 1 || exit 0
	@$(GOLINT) ./... | grep -v vendor/ && exit 1 || exit 0

build:
	docker build -t layer-2a:latest -f ./docker/Dockerfile .

deploy:
	@docker compose up -d || (echo "error deploying layer-2"; exit 1)