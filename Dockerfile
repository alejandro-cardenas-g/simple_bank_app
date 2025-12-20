FROM golang:alpine3.23 AS builder
WORKDIR /app
COPY ./app .
RUN go build -o main main.go

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/main .

COPY app.env app.env

EXPOSE 3000
CMD [ "/app/main" ]