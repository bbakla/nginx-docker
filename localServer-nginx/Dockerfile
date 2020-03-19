FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get "github.com/gorilla/mux"
RUN  go build -o main .

EXPOSE 8091

ENTRYPOINT ["/app/main"]