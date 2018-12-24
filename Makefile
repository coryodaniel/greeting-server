.PHONY: run go.build docker.build docker.push docker.shell clean

BIN_NAME=hello-world
IMAGE_NAME=coryodaniel/${BIN_NAME}
IMAGE_URL=quay.io/${IMAGE_NAME}
CONTAINER_NAME=${BIN_NAME}-dbg

run: go.build
	./$(BIN_NAME)

go.build:
	go test
	go fmt
	go build

docker.build:
	docker build -t $(IMAGE_URL) .

docker.push:
	docker push $(IMAGE_URL)

docker.shell:
	docker run -it $(IMAGE_URL):latest /bin/sh

clean:
	-docker ps --format "{{.Names}}" -a | grep "$(CONTAINER_NAME)" | xargs docker stop
	-docker ps --format "{{.Names}}" -a | grep "$(CONTAINER_NAME)" | xargs docker rm
	rm $(BIN_NAME)
