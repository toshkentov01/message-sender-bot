# workspace (GOPATH) configured at /go
FROM golang:1.17.5 as builder


RUN mkdir /message-sender-bot
ADD . /message-sender-bot/
WORKDIR /message-sender-bot

RUN go mod tidy
RUN go build -o main ./cmd/main.go
CMD ["./main"]