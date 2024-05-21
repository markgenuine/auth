FROM golang:1.22.2-alpine AS builder

COPY . /github.com/markgenuine/auth/source/
WORKDIR /github.com/markgenuine/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/server/main.go

FROM alpine:3.19.1

WORKDIR /root/
COPY --from=builder /github.com/markgenuine/auth/source/bin/auth_server .
COPY local.env .

CMD ["./auth_server", "-config-path", "local.env"]