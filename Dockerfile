FROM gcr.io/cloud-builders/go
WORKDIR /go/src/github.com/priyawadhwa/next-2019-demo
COPY . .
RUN go build -o /hello main.go

FROM scratch
COPY --from=0 /hello /hello
CMD ["/hello"]
EXPOSE 8080
