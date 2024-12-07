FROM golang:alpine

WORKDIR /service
COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o bin/service cmd/service/main.go

EXPOSE 50055

CMD ["sh", "-c", "goose -dir /service/db/migrations up && ./bin/service"]
