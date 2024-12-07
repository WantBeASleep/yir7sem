FROM golang:alpine AS builder

WORKDIR /service
COPY . .

RUN go build -o bin/service cmd/service/main.go

# FROM alpine:latest

# WORKDIR /service

# COPY --from=builder /service/bin/service /service/db /service/bin/goose /service/

EXPOSE 50055

CMD [ "./bin/service" ]
