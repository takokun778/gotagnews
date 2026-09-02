ARG GOLANG_VERSION

FROM golang:1.27.1-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/cmd/app

CMD ["air"]
