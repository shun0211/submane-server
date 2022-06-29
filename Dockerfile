FROM golang:latest AS builder

WORKDIR /submane-server
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api .

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# NOTE: 現在のディレクトリにmigrateのバイナリファイルが作られる
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# NOTE: 実行環境と同じになるように指定する
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO_ENV=development
RUN go build server.go

# NOTE: golangのイメージは大きすぎるが、実行時には不要なので軽量なalpineを使って実行する golangと比べて150分の1になる
FROM alpine:latest AS production
# NOTE: これまでの構築ステージをコピー元として指定するために--form=<名前>フラグを使っている
#       COPY <COPY元> <COPY先>
COPY --from=builder /submane-server/server ./
COPY --from=builder /submane-server/migrate /bin/migrate

ENV GO_ENV=production
# ENV FRONT_URI=http://localhost:3000
# ENV POSTGRES_DATABASE=submane_db
# ENV POSTGRES_HOSTNAME=postgres
# ENV POSTGRES_USERNAME=gorm
# ENV POSTGRES_PASSWORD=gorm
# ENV POSTGRES_PORT=5432

CMD [ "./server" ]
