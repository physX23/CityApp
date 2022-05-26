# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:buster AS build

WORKDIR /app

COPY . ./

RUN go mod download && go mod tidy

RUN go build ./cmd/main

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /docker-cityapp

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/docker-cityapp/main"]
