ENTRYPOINT := cmd/dockerhub-update/main.go
BIN        := dockerhub-update

all: build docker_build

build:
	go build -o $(BIN) $(ENTRYPOINT)

docker_build:
	docker build . -t dockerhub-update
