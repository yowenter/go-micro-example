FROM golang:1.13

ENV GOPATH=/app
ENV  GO111MODULE=on

RUN mkdir -p /app
ADD . $GOPATH/
WORKDIR /app/src/wuu.com/example
EXPOSE 5000
RUN go build server/main.go
CMD ["/app/src/wuu.com/example/main", "--server_address=0.0.0.0:5000", "--registry_address=192.168.1.12:2379","--server_advertise=192.168.1.12:5000"]