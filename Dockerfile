FROM golang:1.23.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o listener ./cmd/listener/main.go

RUN go build -o processor ./cmd/processor/main.go

