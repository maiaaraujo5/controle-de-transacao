FROM golang:1.22.1-alpine as builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -o ./dist/main cmd/main.go

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/configs/ configs
COPY --from=builder /app/dist/main .

EXPOSE 8080

CMD ["./main"]