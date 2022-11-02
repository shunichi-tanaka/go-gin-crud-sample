FROM golang:1.19.1

WORKDIR $GOPATH/go/src
COPY . $GOPATH/go/src

RUN go mod tidy
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-crud-example"]