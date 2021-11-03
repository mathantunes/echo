FROM golang:latest

WORKDIR /server

COPY . .

RUN go mod download

RUN go build ./app/app.go

EXPOSE 7

CMD ["app"]
