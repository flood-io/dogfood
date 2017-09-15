FROM golang:latest
RUN apt-get update && apt-get install upx -y
RUN go get github.com/gorilla/mux
RUN go get github.com/Pallinder/sillyname-go

ADD src /go/src/
ADD config /config
WORKDIR /go/src/

CMD ["go", "run", "main.go"]
