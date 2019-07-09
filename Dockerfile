FROM golang:latest

RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux
RUN go get github.com/globalsign/mgo
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/stretchr/testify/mock

ADD . /go/src/go_songs
WORKDIR /go/src/go_songs

RUN go build

ENTRYPOINT ./go_songs

EXPOSE 8000