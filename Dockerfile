FROM golang:alpine AS builder

WORKDIR /work

RUN apk update && apk add --no-cache \
  git ca-certificates \
  && update-ca-certificates

COPY go.mod go.sum ./

RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags="-w -s" -o /bin/dockerhub-update \
  ./cmd/dockerhub-update

FROM scratch
LABEL maintiner="Charles Kenney <me@ch4z.io>"

COPY --chown=0:0 --from=builder /bin/dockerhub-update /bin/dockerhub-update
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/bin/dockerhub-update"]
