FROM golang:1.17

RUN mkdir /go-elk

ADD . /go-elk

WORKDIR /go-elk

RUN go build -o server .

CMD ["/go-elk/server"]