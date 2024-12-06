FROM golang:alpine AS builder

WORKDIR /service
COPY . .

RUN GOBIN=$(pwd)/bin go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o bin/service cmd/service/main.go

FROM alpine:latest

WORKDIR /service

COPY --from=builder /service/bin/service /service/db /service/bin/goose /service/

EXPOSE 50055

CMD [ "./goose -dir db/migrations up && ./service" ]
