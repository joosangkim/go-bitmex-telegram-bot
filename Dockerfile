FROM golang:1.13-alpine3.11

WORKDIR /go/src/app
COPY main.go .

RUN go build . app
# RUN go get -d -v ./...
# RUN go install -v ./...

CMD ["app"]