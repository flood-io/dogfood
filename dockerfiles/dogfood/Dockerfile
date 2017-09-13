FROM golang:latest
RUN apt-get update && apt-get install upx -y
RUN go get github.com/gorilla/mux

ADD src /go/src/
WORKDIR /go/src/

CMD ["go", "run", "main.go"]
