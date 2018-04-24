FROM golang

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY get.sh /go/src/app
RUN bash get.sh

COPY ./src /go/src/app

RUN go get -v ./...

RUN go build

CMD app
COPY ./static /go/src/app/static

EXPOSE 3004