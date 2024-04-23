FROM golang:1.21.9-bookworm

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy -e
RUN go mod verify
RUN go mod download

RUN go build -o bin/main cmd/main.go

CMD ["/usr/src/app/bin/main"]