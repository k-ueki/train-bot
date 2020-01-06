FROM golang:latest

RUN go build

CMD ["go", "run","main.go"]
