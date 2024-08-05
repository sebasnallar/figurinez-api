FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . ./

EXPOSE 8080

CMD ["CompileDaemon", "--build=go build -o /figurinez-api ./cmd", "--command=/figurinez-api"]
