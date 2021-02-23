FROM golang:alpine as builder
RUN go install github.com/ochko/go-tls-check/cmd/tls-check@v1.0.0

FROM alpine:latest
COPY --from=builder /go/bin/tls-check /usr/local/bin/
ENTRYPOINT ["tls-check"]
