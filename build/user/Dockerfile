FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bootstrap ./cmd/api/user

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=build /bootstrap /bootstrap

ENTRYPOINT ["/bootstrap"]
