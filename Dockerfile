# build stage
FROM golang:alpine as build-env
MAINTAINER mdouchement

RUN apk upgrade

WORKDIR /captive

ENV CGO_ENABLED 0
ENV GO111MODULE on

COPY . .
# Dependencies
RUN go mod download

RUN go build -ldflags "-s -w" -o captive .

# final stage
FROM scratch
MAINTAINER mdouchement

COPY --from=build-env /captive/captive /usr/local/bin/

EXPOSE 8080
CMD ["captive"]
