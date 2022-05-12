FROM golang:latest AS build

WORKDIR /go/src/submane-server/api
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api .

# -oはoutput
RUN go build -o /go/bin/app
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD ["air"]
