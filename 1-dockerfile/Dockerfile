FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /sample-go-app

FROM alpine:3.15.4

WORKDIR /

COPY --from=builder /sample-go-app /sample-go-app

ENTRYPOINT [ "/sample-go-app" ]