ARG GOLANG_VERSION

FROM golang:1.20.5-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/cmd/app

CMD ["air"]
