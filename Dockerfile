FROM golang:latest AS build

WORKDIR /go/src/submane-server/api
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api .
ENV GO_ENV=development

# -oはoutput
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate
RUN go build server.go

FROM alpine:latest
COPY --from=build . .

CMD [ "/server $PORT" ]
