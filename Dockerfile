# syntax=docker/dockerfile:1
FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /boobytrap DiscordBoobyTrap/cmd/discordboobytrap

## Deploy
FROM scratch
COPY --from=build /boobytrap /boobytrap
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/boobytrap"]