FROM golang:1.14
RUN apt-get update -y && apt-get install tree -y
ENV GO111MODULE=on
RUN go get -u golang.org/x/lint/golint
RUN mkdir -p /workspace
WORKDIR /workspace
COPY go.mod ./
RUN go mod download -json
COPY cmd ./cmd
COPY pkg ./pkg
RUN go build ./...
RUN go test ./...
RUN golint -set_exit_status ./...
