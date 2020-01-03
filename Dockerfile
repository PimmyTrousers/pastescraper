FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o pastescrape .

FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/pastescrape /app/

WORKDIR /app
ENTRYPOINT ["./pastescrape", "-config", "./config.yml"]
