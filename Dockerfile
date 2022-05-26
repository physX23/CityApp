# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:buster AS build

WORKDIR /app

COPY . ./

<<<<<<< HEAD
RUN go mod download && go mod tidy
=======
RUN go mod download
>>>>>>> 8ccb34f2f0af774c42f99376706001381862c23f

RUN go build ./cmd/main

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /docker-cityapp

EXPOSE 80

<<<<<<< HEAD
ENTRYPOINT ["/docker-cityapp/main"]
=======
USER nonroot:nonroot

ENTRYPOINT ["/docker-cityapp"]
>>>>>>> 8ccb34f2f0af774c42f99376706001381862c23f
