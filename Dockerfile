FROM golang:latest

RUN mkdir /myserver

COPY . /myserver

WORKDIR /myserver

RUN go build -o main .

CMD ["/myserver/main"]