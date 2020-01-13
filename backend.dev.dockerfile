FROM golang:latest

WORKDIR /app

RUN go get -u github.com/cespare/reflex

EXPOSE 5000

ENTRYPOINT reflex -d none -s -R vendor. -r \.go$ -- go run -v ./cmd/apiserver

