FROM golang:1.24-alpine AS build

RUN apk add --no-cache gcc g++

WORKDIR /build
COPY go.mod /build/go.mod
COPY go.sum /build/go.sum
RUN go mod download && go mod verify
COPY . /build/
RUN CGO_ENABLED=1 go build -v

FROM alpine:3.21 AS run

WORKDIR /usr/local/bin
COPY --from=build /build/borders /usr/local/bin/borders
ENTRYPOINT ["/usr/local/bin/borders"]