FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o pastescrape .

FROM alpine
COPY --from=builder /build/pastescrape /app/

WORKDIR /app
ENTRYPOINT ["./pastescrape", "-config", "./config.yml"]
