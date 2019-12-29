ENTRYPOINT := cmd/dockerhub-update/main.go
BIN        := dockerhub-update
IMG        := charliekenney23/dockerhub-update

all: deps test build docker_build

deps:
	go mod download
	go mod verify

test:
	go test -v ./...

build:
	go build -o $(BIN) $(ENTRYPOINT)

docker_build:
	docker build . -t $(IMG)
