FROM golang:latest 

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app/cmd

RUN go get github.com/psavelis/golang-fluentd-stdout/middlewares

RUN go build -o main . 
CMD ["/app/cmd/main"]