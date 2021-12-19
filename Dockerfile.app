FROM golang:1.17

COPY . /go/src/app

WORKDIR /go/src/app/cmd/app

RUN go build -race -o app main.go

EXPOSE 9090

CMD ["./app"]
