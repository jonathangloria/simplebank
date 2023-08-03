# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./db/migration

EXPOSE 8080
RUN chmod +x /app/start.sh
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]