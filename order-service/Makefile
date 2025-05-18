IMAGE_NAME=rootduck/order-service
TAG=latest
FULL_IMAGE=$(IMAGE_NAME):$(TAG)

.PHONY: build push

build:
	docker build -t $(FULL_IMAGE) .

push:
	docker push $(FULL_IMAGE)

start-server:
	go run main.go server

start-worker:
	go run main.go worker
