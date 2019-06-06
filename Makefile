GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
DOCKER_IMAGE_NAME=docker.henjinet.com:5000/article_manage:latest

.PHONY: dev
dev:
	$(GOBUILD) -o ./main/main ./main  && cd ./main && ./main

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/app_linux -v

.PHONY: build_docker
build_docker: build_linux
	docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: up_docker
up_docker: build_docker
	docker push $(DOCKER_IMAGE_NAME)
