FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /app

COPY . /app/

RUN go build -o golibrary

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/golibrary /app/

RUN chmod +x /app/golibrary

CMD ["/app/golibrary"]