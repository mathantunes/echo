FROM golang:latest

EXPOSE 7

WORKDIR /server

COPY . .

RUN go mod tidy

RUN go build ./app/app.go

CMD ["./app"]
