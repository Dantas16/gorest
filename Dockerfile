FROM golang

ADD . /home/api

WORKDIR /home/api

RUN go get github.com/gorilla/mux go.mongodb.org/mongo-driver

EXPOSE 8080

CMD ["go","run","main.go"]