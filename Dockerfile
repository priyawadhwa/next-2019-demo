FROM golang:1.10
WORKDIR /go/src/github.com/priyawadhwa/next-2019-demo
COPY . .
RUN go build -o /hello main.go
CMD ["/hello"]
EXPOSE 8080
