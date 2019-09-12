FROM golang:1.13-alpine3.10 as builder
ENV GO111MODULE=on
COPY main.go .
RUN go build -o /app main.go

FROM alpine:3.10
ENTRYPOINT ["./app"]
COPY --from=builder /app .
