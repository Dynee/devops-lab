.PHONY: build-app
build-app: 
	go build ./cmd/cat-service/main.go

.PHONY: build-docker
build-docker: 
	docker build --build-arg GO_VERSION=1.18 -t cat-service .
