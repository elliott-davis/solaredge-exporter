FROM golang:1.13-stretch as build
LABEL maintainer="Elliott Davis"

ENV GO111MODULE=on

COPY ./ /go/src/github.com/elliott-davis/solaredge-exporter
WORKDIR /go/src/github.com/elliott-davis/solaredge-exporter

RUN go mod download \
    && go test ./... \
    && CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM alpine:3.10

RUN apk --no-cache add ca-certificates \
     && addgroup exporter \
     && adduser -S -G exporter exporter
USER exporter
COPY --from=build /bin/main /bin/main
ENV LISTEN_PORT=9171
EXPOSE 9171
ENTRYPOINT [ "/bin/main" ]
