FROM golang:1.20 as builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
