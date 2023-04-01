# syntax=docker/dockerfile:1
FROM golang:1.20-alpine as build
ARG VERSION=dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -ldflags="-X 'main.Version=$VERSION'" -o /boobytrap DiscordBoobyTrap/cmd/discordboobytrap

## Deploy
FROM scratch
COPY --from=build /boobytrap /boobytrap
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/boobytrap"]