FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8080

ENTRYPOINT [ "/app/music-api" ]