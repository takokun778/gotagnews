ARG GOLANG_VERSION

FROM golang:1.26.0-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/cmd/app

CMD ["air"]
