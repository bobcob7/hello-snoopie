FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o hello-snoopie .

FROM alpine:latest
COPY --from=builder /app/hello-snoopie .
EXPOSE 80
CMD ["./hello-snoopie"]
