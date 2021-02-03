FROM golang:1.15 AS build

WORKDIR /build
COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o base github.com/buzyakabarbuzyaka/base/cmd

FROM debian:stretch AS use
WORKDIR /app
COPY --from=build /build/base .
RUN mkdir -p /app/log
COPY deploy/prod-conf.yaml conf.yaml
RUN apt-get update && apt-get install ca-certificates -y

CMD ["./base"]