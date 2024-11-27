FROM golang:1.23.0 as builder

WORKDIR /app

COPY . .

RUN go install github.com/go-task/task/v3/cmd/task@latest

RUN task build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/cmd ./med

EXPOSE 50055

CMD ["./med"]
