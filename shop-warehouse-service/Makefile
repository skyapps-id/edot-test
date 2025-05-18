IMAGE_NAME=rootduck/shop-warehouse-service
TAG=latest
FULL_IMAGE=$(IMAGE_NAME):$(TAG)

.PHONY: build push

build:
	docker build -t $(FULL_IMAGE) .

push:
	docker push $(FULL_IMAGE)

start:
	go run main.go
	