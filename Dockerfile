FROM golang:latest

WORKDIR /train-bot

COPY ./ /train-bot

RUN go build

CMD ["go", "run","main.go"]
