FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o pastescrape .

FROM alpine
ARG config_file=example.yml
ENV config_file=${config_file}

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/pastescrape /app/
COPY --from=builder /build/*.yml /app/
RUN ls /app
WORKDIR /app
ENTRYPOINT ./pastescrape -config ${config_file}