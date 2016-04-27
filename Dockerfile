FROM golang:1.6

ENV GO15VENDOREXPERIMENT=1
ENV GOBIN $GOPATH/bin

RUN go get -u github.com/tools/godep

WORKDIR /go/src/github.com/NOX73/tech-ops-challenge

ADD . /go/src/github.com/NOX73/tech-ops-challenge

RUN godep restore

CMD go run ./main.go
