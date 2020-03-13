FROM golang:1.14-alpine

ADD . /home/api

WORKDIR /home/api

EXPOSE 8080

CMD ["go","run","main.go"]