FROM golang:alpine

WORKDIR /service

COPY . .

RUN go build -o bin/service cmd/service/main.go

EXPOSE 8080

CMD [ "./bin/service" ]
