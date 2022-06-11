
deps:
	go mod tidy
	go mod vendor

build:
	docker build -t auth:latest -f ./docker/Dockerfile .