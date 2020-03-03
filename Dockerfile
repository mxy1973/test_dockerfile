FROM golang:latest

WORKDIR $GOPATH/src/github.com/go_test
ADD . $GOPATH/src/github.com/go_test
#RUN go build .

EXPOSE 8000
#ENTRYPOINT ["./go_test"]
