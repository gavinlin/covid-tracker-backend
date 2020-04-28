FROM golang:alpine AS build

ENV GO11MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build

RUN apk --update upgrade
RUN apk add --no-cache git make build-base
RUN apk add --update gcc=9.2.0-r4 g++=9.2.0-r4
RUN apk add sqlite

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:edge
WORKDIR /dist

RUN apk --update upgrade
RUN apk add sqlite

COPY --from=build /build/main /dist/main

EXPOSE 8080

CMD ["/dist/main"]