# Build image
FROM golang:1.9-alpine AS build-env

ENV APP=flyte-ticker

RUN apk add --no-cache git curl
RUN git config --global http.sslVerify false
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR $GOPATH/src/github.com/HotelsDotCom/$APP/

COPY . ./

RUN dep ensure
RUN go test ./...
RUN CGO_ENABLED=0 go build

# Run image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
ENV APP=flyte-ticker
COPY --from=build-env /go/src/github.com/HotelsDotCom/$APP/$APP /app/$APP
ENTRYPOINT "/app/$APP"
