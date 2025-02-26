FROM golang:1.23.5-alpine AS builder

WORKDIR /app
COPY . .

# Build app
WORKDIR /app/app
RUN go build -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app/main /main
COPY --from=builder /app/app/.env /.env

ENTRYPOINT ["/main"]
EXPOSE 8080