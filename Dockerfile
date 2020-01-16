FROM golang:latest

WORKDIR /go/src/github.com/k-ueki/train-bot

COPY ./ /go/src/github.com/k-ueki/train-bot

RUN go build

EXPOSE 8888

CMD ["go", "run","main.go"]
