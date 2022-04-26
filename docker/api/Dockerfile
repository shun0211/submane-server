FROM golang:latest AS build

WORKDIR /go/src/submane-server/api
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api .

RUN go build -o /go/bin/app
# RUN go build -o /go/bin/app && \
  # go get gopkg.in/urfave/cli.v2@master && \
  # go get github.com/oxequa/realize
