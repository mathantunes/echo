FROM golang:latest

WORKDIR /server

COPY . .

RUN go mod tidy

RUN go build ./app/app.go

CMD ["./app"]
