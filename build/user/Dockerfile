FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bootstrap .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=build /bootstrap /bootstrap

ENTRYPOINT ["/bootstrap"]
