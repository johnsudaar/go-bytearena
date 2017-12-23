FROM golang
RUN mkdir -p $GOPATH/src/myagent
COPY . $GOPATH/src/myagent
RUN go get myagent/...
RUN go build -o /go/agent myagent
ENTRYPOINT ["/go/agent"]
